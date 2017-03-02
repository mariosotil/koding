package app

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"koding/klient/uploader"
	"koding/logrotate"
)

// UploadedFile describes a file that was uploaded to
// Koding log storage.
type UploadedFile struct {
	Local  string `json:"local"`           // local name of the file
	Remote string `json:"remote"`          // remote name of the file in Koding log storage
	Error  string `json:"error,omitempty"` // if upload failed, text message of the failure
}

func (u *UploadedFile) err() string {
	if u.Error != "" {
		return u.Error
	}
	return "-"
}

type uploadCmd struct {
	up   *uploader.Uploader
	all  bool
	file string
	json bool
	out  io.Writer
}

func (cmd *uploadCmd) registerFlags(f *flag.FlagSet) {
	f.BoolVar(&cmd.all, "all", false, "Uploads all built-in logs.")
	f.StringVar(&cmd.file, "f", "", "Upload the specified log/text file.")
}

func (cmd *uploadCmd) uploadFile(file string) (*UploadedFile, error) {
	u, err := cmd.up.UploadFile(cmd.file, 0)
	if err != nil && !logrotate.IsNop(err) {
		return nil, err
	}

	return &UploadedFile{
		Local:  cmd.file,
		Remote: u.String(),
	}, nil
}

func (cmd *uploadCmd) upload() (uploaded []UploadedFile, err error) {
	if !cmd.all {
		if cmd.file == "" {
			return nil, errors.New("invalid empty value for -f flag")
		}

		u, err := cmd.uploadFile(cmd.file)
		if err != nil {
			return nil, err
		}

		return []UploadedFile{*u}, nil
	}

	for _, file := range uploader.LogFiles {
		u, e := cmd.uploadFile(file)

		if e == nil {
			uploaded = append(uploaded, *u)
			continue
		}

		if err == nil {
			err = e
		}

		uploaded = append(uploaded, UploadedFile{
			Local: file,
			Error: e.Error(),
		})
	}

	return uploaded, err
}

func (cmd *uploadCmd) stdout() io.Writer {
	if cmd.out != nil {
		return cmd.out
	}
	return os.Stdout
}

// Upload implements a handler for the following internal command:
//
//   klient upload [--all] [--json] [-f <file>]
//
func (k *Klient) Upload(args ...string) error {
	f := flag.NewFlagSet("klient-upload", flag.ContinueOnError)
	cmd := &uploadCmd{
		up: k.uploader,
	}

	cmd.registerFlags(f)

	if err := f.Parse(args); err != nil {
		return err
	}

	uploaded, err := cmd.upload()
	if err != nil {
		return err
	}

	if cmd.json {
		enc := json.NewEncoder(cmd.stdout())
		enc.SetIndent("\t", "")
		return enc.Encode(uploaded)
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	fmt.Fprintln(w, "LOCAL\tREMOTE\tERROR")

	for _, u := range uploaded {
		fmt.Fprintf(cmd.stdout(), "%s\t%s\t%s\n", u.Local, u.Remote, u.err())
	}

	return w.Flush()
}
