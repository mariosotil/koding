package machinegroup

import (
	"errors"
	"net"
	"strconv"
	"sync"
	"time"

	"koding/kites/tunnelproxy/discover"
	"koding/klient/machine"
	"koding/klient/machine/client"
)

// DynamicSSHFunc is an adapter that allows to retrieve information needed to
// make SSH connection to remote machine.
type DynamicSSHFunc func() (username, host string, port int, err error)

// SSHRequest defines machine group ssh info request.
type SSHRequest struct {
	// ID is a unique identifier for the remote machine.
	ID machine.ID `json:"id"`

	// Username defines the remote machine user name. This field is optional, if
	// not set, the current remote user will be used instead.
	Username string `json:"username"`

	// PublicKey contains local machine public key content which is meant to be
	// added to remote machine authorized_keys file. This field is optional, if
	// not set, no keys will be added.
	PublicKey string `json:"public_key"`
}

// SSHResponse defines machine group ssh info response.
type SSHResponse struct {
	// Username defines the remote machine user name.
	Username string `json:"username"`

	// Host is the remote machine host used for SSH connections.
	Host string `json:"host"`

	// Port defines remote machine SSH listening port.
	Port int `json:"port"`
}

// SSH prepares remote machine for SSH connection and returns necessary data
// to connect it via SSH protocol.
func (g *Group) SSH(req *SSHRequest) (*SSHResponse, error) {
	if req == nil {
		return nil, errors.New("invalid nil request")
	}

	username, err := g.ensureSSHPubKey(req.ID, req.Username, req.PublicKey)
	if err != nil {
		return nil, err
	}

	username, host, port, err := g.dynamicSSH(req.ID, username)()
	if err != nil {
		return nil, err
	}

	return &SSHResponse{
		Username: username,
		Host:     host,
		Port:     port,
	}, nil
}

func (g *Group) ensureSSHPubKey(id machine.ID, username, pubKey string) (string, error) {
	// How long to wait for a valid client.
	const timeout = 30 * time.Second

	var err error

	dynClient := func() (client.Client, error) { return g.client.Client(id) }
	// Use provided username or ask remote machine to return it.
	c := client.NewSupervised(dynClient, timeout)
	if username == "" {
		if username, err = c.CurrentUser(); err != nil {
			return "", err
		}
	}

	// Add pubic key to remote machine authorized keys.
	if pubKey != "" {
		if err := c.SSHAddKeys(username, pubKey); err != nil {
			return "", err
		}
	}

	return username, nil
}

func (g *Group) dynamicSSH(id machine.ID, username string) DynamicSSHFunc {
	var (
		mtx  sync.Mutex
		addr machine.Addr
		host string
		port int
	)

	return func() (string, string, int, error) {
		// Check for tunneled connections.
		a, err := g.address.Latest(id, "tunnel")
		if err != nil {
			// There are no tunnel addresses. Get latest IP.
			if a, err = g.address.Latest(id, "ip"); err != nil {
				return "", "", 0, err
			}

			return username, a.Value, 0, nil
		}

		// Use a pseudo-cache in order to not call discover each time.
		mtx.Lock()
		if addr == a {
			mtx.Unlock()
			return username, host, port, nil
		}
		mtx.Unlock()

		// Discover tunnel SSH address.
		endpoints, err := g.discover.Discover(a.Value, "ssh")
		if err != nil {
			return "", "", 0, err
		}

		tunnelAddr := endpoints[0].Addr
		// We prefer local routes to use first, if there's none, we use first
		// discovered route.
		if e := endpoints.Filter(discover.ByLocal(true)); len(e) != 0 {
			// All local routes will do, typically there's only one,
			// we use the first one and ignore the rest.
			tunnelAddr = e[0].Addr
		}

		h, p, err := net.SplitHostPort(tunnelAddr)
		if err != nil {
			h, p = tunnelAddr, "0"
		}

		n, err := strconv.Atoi(p)
		if err != nil {
			return "", "", 0, err
		}

		// Cache results.
		mtx.Lock()
		addr, host, port = a, h, n
		mtx.Unlock()

		return username, host, port, nil
	}
}
