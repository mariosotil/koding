// Package testkeys contains RSA keys and TLS certificates for using in examples and tests.
package testkeys

///////////////////////////////////////////////////////////////
//
// RSA Key
//
// RSA key is generated with following commands:
//     openssl genrsa -out testkey.pem 2048
//     openssl rsa -in testkey.pem -pubout > testkey_pub.pem
//

const Public = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxqvspbT/HInFTJfyJ14f
WB6sSeMC8f11iwHNOwj17Pm4Qpakv+RvdMEm4H/nkEBJiX9l5U4ooKqeLFGSMj8P
vd9u83ef36guedQTp6As5KcTp8Ik+oYtsGuG812aZ21enaOdL3KaKSMPl5DVz2n8
du3JLkV2BdhGFs4ylio/H/W5uVL52ceZ7e2kpfvR4NKjct5qmOLIF3sVaotRvKSm
CDnQLMPX3dcM53oA9zPOY7ok9x+Ga2DbPgFSUAt2jr0MyUVecYNX9yMTEzCtfdBN
UgqZFjni30bOLIvW9Cf/YfM71Q15WayUjfJhb7jR6lC65WT3dKentuORAmnerRr2
vwIDAQAB
-----END PUBLIC KEY-----`

const Private = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxqvspbT/HInFTJfyJ14fWB6sSeMC8f11iwHNOwj17Pm4Qpak
v+RvdMEm4H/nkEBJiX9l5U4ooKqeLFGSMj8Pvd9u83ef36guedQTp6As5KcTp8Ik
+oYtsGuG812aZ21enaOdL3KaKSMPl5DVz2n8du3JLkV2BdhGFs4ylio/H/W5uVL5
2ceZ7e2kpfvR4NKjct5qmOLIF3sVaotRvKSmCDnQLMPX3dcM53oA9zPOY7ok9x+G
a2DbPgFSUAt2jr0MyUVecYNX9yMTEzCtfdBNUgqZFjni30bOLIvW9Cf/YfM71Q15
WayUjfJhb7jR6lC65WT3dKentuORAmnerRr2vwIDAQABAoIBAQCiQT9ghlXj5NDy
IRZv+Hr46PQk0/ZP1ITvJmWK1WKkS76lA4tB/TrZy/YIW+2u7hIg7Z82i4K61R2h
bG1OBBrDI3vl88jqTNzIOuBignqns5xl4jV5NaSS9P6eO9AisDwZ2spBOgOVnUDa
tDsqzpCWQNWRqMEfPObhMT43TgGIAyCsDqwBZT4ZFZHEYpJKgajGJQJPam6e44MU
eLbm17EVMQiD175jsDpjYYvVFLD8VBW9E+EJgk/pO809XWPNW5A7Q/YZexhPaLQw
gipcb6K3vimkMPqlraR/NUBUz9hzeItD94weGhMBOJfIjE8pEbFz1ef3M+86zjcs
G/u4utghAoGBAP7pFNFvtXefr3O2iGWFjEoaKrgGjxH1rx7VXNjkSVEi7FGyWjcE
sJUqjxONGWAjLLviYfz8A1knnVYQRi5PTnPqBUqNFmF3tcFVsqJ8uPFw0FDIGwSQ
DCZYVxFwEHnSttOi0dhxnCMEEiPDd2i33x6Ayq5D7mfDzR1R1UiGSuh1AoGBAMeF
TrTXCdo5yveM4hseXCmYTlYMkwvkhvIzoWYKeBsdMfSclJZUHZ52UodfgHgBIEE5
bVO52QPXFeDY2WJKuQLxe1dEU5/DIQFX6VG9kPajaMP/tOxR+5F6qXAywzjJo9ao
36mafoV/B74spbM/t5ApG0DEYb/Uy6NMFg0IMJvjAoGAHCEcawAoxkqY0ewSoSFb
mdHJZQVURGcYLA7fcA/BB3MTIPDNg8TAt7U6lK29Xx6CkTkqFwC4Xd5BPn4NQI7/
2Uq5ysBG9/p91Bfg3o1k2z/XvRJX6Oj8bj7RXNtA9jCKfFA7V+cxk6ufVNL9GlEw
De2ocBKO/GRe/h1Jq/clzQECgYEAvXS5OwEiquTL9WsC3QtkKl8kq5L/IoDbHyul
q7ZvPBfp8sOoWGMBSRJD5aXq4Ij5LtYwGRvVlU7syHhBW5NmwQO6wi/wVDk+ov7u
X57pAzpB5jNugiSLBCZ8tDyPoeHewJvU3kEEuRDZslzE70Oms65E4ahYxYZz49kA
61kGh08CgYAiYrKMQ3iic1MI9phZd2XeG02t2bxTlP76fIXUMe5rJirnjNLGOIyd
Bt+/AzX98FU9DGJIZTm38OzCFJnbp/awI1QnUqPTtLV3WYtdH58kQUByszdyt7aQ
RBO19jSxMDjWlQeZbfh/goPdXvuCu8/mgrxa8xO3anx7M5WGrSKS1A==
-----END RSA PRIVATE KEY-----`

//////////////////////////////////////////////////

const PublicSecond = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyk80fw95zWJ/WZaaYD8C
l65QC9sHwE1SoJvMgV7nauxWL63aj5MIFgYbrgpG9qr2FBTOvx+tHFkuLzOkykoO
/blYXqFhsq4ok3euVX0m61tz1IfiEkw+w+fxJDQzzBF0kPrwTlFwND9bx/n89dak
CtfcyOxiapM5P27oiUzgFNyHuts9+SeuN92Y8zKqk6XKein6emnRU7OeUJALsUTx
akvGCKwkwDArWYNUF5Vsg7TaIRDgkDZ1frtWI3D34fFyjcEPme7xkXzgsocHM1gL
4M+YqqToBzg6/PTJRRWXVlK1xfCFl9K9WNG475msjojCxOKOLrFlbaVg4LAO2xT7
yQIDAQAB
-----END PUBLIC KEY-----`

const PrivateSecond = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAyk80fw95zWJ/WZaaYD8Cl65QC9sHwE1SoJvMgV7nauxWL63a
j5MIFgYbrgpG9qr2FBTOvx+tHFkuLzOkykoO/blYXqFhsq4ok3euVX0m61tz1Ifi
Ekw+w+fxJDQzzBF0kPrwTlFwND9bx/n89dakCtfcyOxiapM5P27oiUzgFNyHuts9
+SeuN92Y8zKqk6XKein6emnRU7OeUJALsUTxakvGCKwkwDArWYNUF5Vsg7TaIRDg
kDZ1frtWI3D34fFyjcEPme7xkXzgsocHM1gL4M+YqqToBzg6/PTJRRWXVlK1xfCF
l9K9WNG475msjojCxOKOLrFlbaVg4LAO2xT7yQIDAQABAoIBADapHcryuHsYkMX4
3e8BN0caLsB1RmvbuGZykdemd6o4/rRVKcc+96FTtyjX2AKPgHs+f/m9qj0Nj1/r
eSu6xMAi0tCGk+n+CjKF1JF+hgRzKiGTMS62cQLnaQzaGeCaGr+NPV47vLAxKjAm
yAT4IExZtGqJC7I14vLTmXp3Tdf0nlp3KkQ3Uamm3gZW3OKNQd3+GbdeHOgWbO7W
eaRilR2uCPkXYdFjho3Q04NzURwrfnlAfKkq1Bb4/vgpoVg35xquHQ7gXXFTr64H
VG11K6pHDImn5QiiCti7fSYx7lbuuMtPBMvh+5EIefXzib3OASpBrmscpzO2rk9v
wV3c4e0CgYEA9SOXLp4QxiAUMlMDMOK23K/+ovO8sgiDXdBHsWpQ8RLjnLsR7mVr
eI/WGTzLtSK4YqWcDquLAyiQids0NmU9g4XLhap1rsf++r231e9bGhzBMP9bYQF1
YcBB9qK5rujCezZLGKb2Javwv7zvszbjS5usNodBbVPuI9P2dXnaJ/sCgYEA00XV
WQ4QyvYbft45dlhvwQ3urqP79JGXjQkeFhU3Q95wNGtOuxBV+7TwZRwulB/RNwFV
eifCmpqQjRX9F4MEag52PmjH63t6GIhvcqkCQ1Qzx8mGs6ydXh6qUs4Xi3t4fk8N
cjgkAZ69GrRP6GhZN3xyWJFfxarMkCK2yGMIjAsCgYEA0YPKufg04/EU8fILPyP2
IGZ3XzSsqQknpe3W6KayaWi4iwNEDxo1oYRl+4n/nWAAcaeT2uH43Qk1h+2HEZqz
2Y5n5WVMUcbzgcDWt41sssOyxrrpkd5aQeK9PhvUUc70MbS0uGwy4v2ytV25DNYz
rDJwHOa7H8LlPU/zTHKJ5zMCgYEAgo3awsdQVTszznggZiNMG64yWjT3UzBMyFhk
AR1nI0dnat0Mr8fuejZbfv+lQN9Qd38ZhPzg4oy02ppF5auOpML/Cp3RPJD26AYX
aHFL9rMntEOyO4FlVW35rmWwYv8PfG35TyWmCmI/VSsrXeBtkT4ToutilVFwS3lI
HhgkhfUCgYEA741+vMOgrhcpsiSefLMIdOr4Iv6NlfHMbWKvWLDiij2O9P/P/UE+
vldcz+pdoqFuADPkJaDP2BuSi3qrMtQTtJs8tD3GYwKXRliA0GWiM+PrQ8TU13NA
CLREiOYwWNgwzHuKvCqCvYoLKP4Ks1eceViCW9v4RyzoMBn9iSiwvzs=
-----END RSA PRIVATE KEY-----`

const PrivateThird = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA3HUeWJZWWfAL8BnCPjiI2ev6tguMEHfyVeBNYA59ivYX/hEV
S3SEpItWGu0BDEFCnRxebJAkjBSjkvCbHDxVbP/XS5LIAVC/A5HnPGLO0zr40I0z
w0atyUGVOoiUWrdseH8BcbtJls9bbBlkI85oi5jl64HLILypM2IL22H3R9dCNsiS
Nznzpb6cEY2840jW2/FfnZPFimwqQU7TLnq+HANzUJycFr4RaS3ZezEiWndADsTf
dId5Lr6ComqTuBEjyaSy9n9Bj1UT0pPmanHDQFlWPjxAFIYGiUSp8zpbPE12qXqX
mFJekOriXHpdLI4jwhsTaoH/4xZNmKDAagHHmQIDAQABAoIBABHcs26+D5UN7Amy
LJjZ/8yYYCHr4ZdcAJiQqfREeEuQEPWW8MDCWhTy7TEVTuZtSzZrnALz0uI+sdow
AIFTKiDAwfLvm6Dvp+pkIfY0k0luF8beCfmiemY6GBs8Abkv72v77mQGprm/Z0rk
68Yy7SmSY96nD4xBfPwaz5cg4uZBFig+t4isrucuSyF/CL32UYtS8K5mmPAVC1Ry
YAreOvo5M1s/aWZVo22iqCcLE+C+r7zJAB6zXYgk3DwcyRUrWv9PC0v+ydkBFiu2
YtlIL2VdnAG0G/C2mOrz53eoR0PPReiimpx94uUCyGREM8+6KFTUpQUJXH1iIZvc
v0Bh9jkCgYEA9pgDRqacvKBNki80oJ0XcPf5lBNQ4umfdJpUXLNnBKZyEt/KPM1T
uPDOcaJDixDKlxDhnX24QPpFNbkDT2VZBmX7KtvYgIZtzMCXF656Z0oKzb0nax9F
+L0fthFst/7WQigZCI6fc5HNZCBvKLQ+8ldc0JubVgBgfAvHKycpObMCgYEA5N3i
hYs0O13mLRaZzErd0vX5uSfAv2bBKtRPlGh0WLnsIUNGDBZDBFQ126oarAbsXyh4
eZ0pdY2MeFng5MVwqH53K5Xf1+WRZ9Cu3yRXXNoBtxvt5QVTpTOQu2ZlQ8c5/aVS
XERlU/To+NUiDZ2HmNR4jiEIuS5GZ0srZDlvO4MCgYBa+UXsILeeVz2G+uddgdcO
FE39NQsYS7xKFk2dTYpqCsX29Jz1xYJnrl9bSNPGUFGNXer+YU1Cm5JA+ZqhBfDM
DX+WFjlPYUjMyxhy4inCRWyq7mfGe2kjFRLH5rHHR7rf+U8S+xsUVewCTnzz9Vtr
qJNTWrT6C3WNDgNShkYNRQKBgQDGLXBR7Rk8cvG2h0SEBywP2dj0wHyCAIsx+Ag0
3UIp9kyG/QbxvnEnlXfRj3lrDs+tEpFH1ZQp63kAWeuIjYMhE+lDP5NGEs6XE7fp
F0Hi8A/eK/wSfUtR2NXxwBfd2ezYO0iXBs6rJX56wausGIVrDEib4fT4zR4/oXKg
Za9w2wKBgBGa+sxBfZTTmf9CMlDNcRmyc5QnmZ5FS1X6yvIxmH+DEzXSomdVJ5r4
vFNKSNxPzkA36INEQiWIOvgq5WdUkECASp9qr/L6Ik4YXP8SUfrQm+zPgdHrmYPc
rFssH67ZKgUd4Q+53i4rUWMxHmMyPZKF0RvFnz3U5h5n99sX5muc
-----END RSA PRIVATE KEY-----`

const PublicThird = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3HUeWJZWWfAL8BnCPjiI
2ev6tguMEHfyVeBNYA59ivYX/hEVS3SEpItWGu0BDEFCnRxebJAkjBSjkvCbHDxV
bP/XS5LIAVC/A5HnPGLO0zr40I0zw0atyUGVOoiUWrdseH8BcbtJls9bbBlkI85o
i5jl64HLILypM2IL22H3R9dCNsiSNznzpb6cEY2840jW2/FfnZPFimwqQU7TLnq+
HANzUJycFr4RaS3ZezEiWndADsTfdId5Lr6ComqTuBEjyaSy9n9Bj1UT0pPmanHD
QFlWPjxAFIYGiUSp8zpbPE12qXqXmFJekOriXHpdLI4jwhsTaoH/4xZNmKDAagHH
mQIDAQAB
-----END PUBLIC KEY-----`

const PublicEvil = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs4KnPdBHaS1RIjECBiu2
V6iRkO7Nf1q6icMLUHBfMhJiWQbS7knFNSfCJ59Zb7girQixJLvMELHn6rgVg1nb
RiUChBCxGul9NU/Gq5Xk0GzP4CiMmKFwd3DZeai3Tsod9txu/3nmJA6Cl40hXTXc
2gNSceBCsNz6WBTpp5QHOBntV516ObtN2GUAyuk5XqiCTwGXCL/Sn6QyEEkTdhaT
g2f//HxPWRR9Wy5juXQxblmJUa3yMUTQ255OW0a8mQa/Lg015SwXhDrPsgkh/5qV
z/J79rCi0ej9Bkh17+0V0EctML50R1y94JPn+KQ7IbzItahz0j8zmtbDcpAwkdmO
lQIDAQAB
-----END PUBLIC KEY-----`

const PrivateEvil = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAs4KnPdBHaS1RIjECBiu2V6iRkO7Nf1q6icMLUHBfMhJiWQbS
7knFNSfCJ59Zb7girQixJLvMELHn6rgVg1nbRiUChBCxGul9NU/Gq5Xk0GzP4CiM
mKFwd3DZeai3Tsod9txu/3nmJA6Cl40hXTXc2gNSceBCsNz6WBTpp5QHOBntV516
ObtN2GUAyuk5XqiCTwGXCL/Sn6QyEEkTdhaTg2f//HxPWRR9Wy5juXQxblmJUa3y
MUTQ255OW0a8mQa/Lg015SwXhDrPsgkh/5qVz/J79rCi0ej9Bkh17+0V0EctML50
R1y94JPn+KQ7IbzItahz0j8zmtbDcpAwkdmOlQIDAQABAoIBAFsL8XdQpGecLIKD
CNvIX/ul6+7usBvgEKy+2IY7+IyU9nzhESr7D6MeP0OJdvtLEYth1TckaSQul8pd
A8xTTvwM2XHSZYGY24CmrcVpiVyNVAIFjwn7F+f8vNEP2amEqh4DP+kkEq5HDcWA
N2PnZdTNyosni6vY6MC0Gq58Tg0Nf6Wewru1upd/wz6xOshRtlIxoLvrmDXqHuA1
QdUmCmMcrvlIvHwvfusgLxzih+rBbXdHTkTUdwElvuHr3AXLv5ZY/F7hi7nv8J6X
NekGhH8uQ+U8BMzXJSS5pupvvrGpaWQlERAj4XZ1XHaN/iXyxv+x+v+Avglg6StM
hxTqWQECgYEA4j/d+ACh+NyiAGED6i8Lt4I68nUf5AMCIKzFW7oyzeXc5BZmC0da
DP6xLrgj5B4YkrWCV52k9e8nQaslG1W8238y61xD1cbQ0YvtB43OgMsyyiD/LNu/
Z86Hz/5nfslDPhF+KzLzF0L0RZJ6mKKriyphn+oqNzqFItFFHYcrza0CgYEAyx1t
vwP8Q5GGC7NdcT7rW85BSnTYxMPshUORQbo6KABwSp3hHImoPUOqiWAPoizBLWDM
/GKu2+2JmD4wTz5sShEeyIiNulmyXicJB8YsF+6bz/93l0mOLkuKF9VgDyk1ocFt
WIEe5wtdURodJx4ky4q28Rl3ZmUh+zrJq3XAEYkCgYBN1Y71TLJsPOr2mmmQXRL4
1LKWyrhn5qkKuKVEwy/LKbLuPM5qPue55Lzrx6mBRuFJR2xJ3A/uE5I7wzcGyl4o
XQAVfC5SEw2vqSWoHZ7XLBCS/PsMYaTdf221nl3YfkDFz5rKHcMHU59Zd+T5Ma02
OSRQsWxIh7dZnQjb+a6WGQKBgFeiVtt3aLvuaZtaxBI8R2fQ0bLCP1SGA+JriJyH
MNhZeBl5jMq3SfNE4qtq2tPp418kyMyL903Eav1Yt5c5I5fBUzrKT/v6/05IIUlN
Y3Df7jIL0xlfDw1CYk5uLYfdC9rCjd8FtsOQz65SSgm6o71+F/hmOHHhaIvwjVqA
72GhAoGAO1FaFhO6h+2tKX7kBDBee5JRqbbtb7yZQyDFRW0iyR9SJEmkMo/K/vAZ
yRlRnuZkz8QRnRJyKI+YRTeQdLVjGnDGBHwJIxDolqgOGqI4RiuNVrjNEnnReptT
56sJylfJpR/5/xWF5O/I+RNJ2q/uirkeprWcWOPOqVMZV4hLJcU=
-----END RSA PRIVATE KEY-----`

////////////////////////////////////////////////////////////////////////////////
//
// TLS certificate
//
// Generated with following command:
//     go run generate_cert.go -ca -duration=87600h0m0s -host="localhost,127.0.0.1"
//
// generate_cert.go is available at http://golang.org/src/pkg/crypto/tls/generate_cert.go
//

const Key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAoSJrA4+/gjNCCqnTq56dLrs9U3oB1iYds5HgUMTVjUtwlRr2
eHsOUZmEHIYLu8dtpqN02x0Zdp4G9nAjhOg4kkJ55MBbYEtFJ+ovLBClK44idXlV
6f+wvPLzqb6R2O1szByYkrwIitTYGjcuQvfIOc2msoRTtT86V3NHNB7i+19DeHzK
jZN0E2j+d/kPZLlLNDau22EMu2MhUM6v5kgbLlb9ly9qZIq+FEWOlvP1k4sQsMBk
9cY6JlV+tXWu2Udm4bNc4foEwhOgptMz1HrBXeq/YJmRd0oxcOsbzqqBaN5+SZat
d/CQ1dR+ymOsSrQC02R+gO6iYL1VJwnS9vLbgQIDAQABAoIBAQCQWz5WJX0jjvpW
hUjd0q7sw9AwLfjhOqqXSlQU6BKNkA1fd5Vh4JS2Y/sodqqiYwsIJB43gv1h+Dcy
psobctPjrGx5lB8IyMY++R2js5HV0HNsnPRsO6bB0EKdSt5s4l/7CI9jvdNFVWP5
MIBgfI3Dw0rZWZWIpgmFeQBMYHNnUMgrjrMsWONFy77+ce5IzeQapqjKGJj8bSPx
QygDmXYsXL11OimZCgGYgHkDhYMBAFCTTqxjIC1et6Vt4Q9tHPq3paYZGmJGxlCM
zW+ACqPhFqlhRBpwasNjPq2DNd5qrxrxQy9lH1k64xj+ufAVOBh2/3fyYiaQdVDf
doLqBza1AoGBAM896ptSbnVw/Ho/V9cJJb0b3uyIWtGnPkEK1qBVm+ggjdkvdZgj
0r2Ij3c5ivcOWyY0rw37p4JMLXMcWGhpT7UIkbyw26/hQ8Hqg1CDjja6rQqk9w+M
fmFKNesvAg+udj5WmnBD9tZbpg3/+X75IwNV0agFGc4xFFcsD6MAmmMzAoGBAMcL
eQk8cY6UMk3x16UvV2pSWzBKGOsPcXYUPZU2/pczFC1z5ZXXacF108KYEwErRRdF
E237dv8wk0vzxy+osj3wPSBopdYCSpT4JsLkQjcfq3whFe1sluhI6uU+wWG+4gOg
70Tcz9+XF6i6dgq8vBvOIf8MF590qSOP1SVXxgZ7AoGBAMFPJOZK5sPfZ3J0YXAw
lSYnuDH0IprLILO082xNTocgzo+WyF7ok2u91OzGHQzENFuX6u3lHmPYwNBN0V1b
VkEsRBy478LIIdbg9CoG/IGqcTyrkdTMHRqNCWxdxdNwzdTqHVa43qcmh+cpLMW2
iGVVCKJsV5zQhQOnEtLd/iSnAoGAEctpx7SapleDY8qCL+dalkvEkT2emq0nxJms
o8Fl41iSrjmVRd8bOktxYg08bbdTlu/6+7MmgUvzby/dF9qqDLWEkXgpk0djKYxB
0tybOthe+Vrv3ej+WMfkEBibK0ToadCic3KWNvQIDnKVCQK0gnvcODP5jKzeTiqW
sJ73Oi8CgYEAsQH3wJXbNr4pJBv72Mmtl0a8SKP16pgMsWznhf2Fcj6T5sUZUSif
NHwIB/4pDzSytQh2CA/YAViy/V+M/R/DACCqc/1oAbuoD5jeSiWFokZS+WI74+/H
Hiee/RxoZ9I9UTE8AonPfRgp9A86gwox7qhTeqNCUbWVtVDQW4ZxQ3Y=
-----END RSA PRIVATE KEY-----`

const Cert = `-----BEGIN CERTIFICATE-----
MIIC7zCCAdmgAwIBAgIBADALBgkqhkiG9w0BAQUwEjEQMA4GA1UEChMHQWNtZSBD
bzAeFw0xNDAyMDcwNTAxMjZaFw0yNDAyMDUwNTAxMjZaMBIxEDAOBgNVBAoTB0Fj
bWUgQ28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQChImsDj7+CM0IK
qdOrnp0uuz1TegHWJh2zkeBQxNWNS3CVGvZ4ew5RmYQchgu7x22mo3TbHRl2ngb2
cCOE6DiSQnnkwFtgS0Un6i8sEKUrjiJ1eVXp/7C88vOpvpHY7WzMHJiSvAiK1Nga
Ny5C98g5zaayhFO1PzpXc0c0HuL7X0N4fMqNk3QTaP53+Q9kuUs0Nq7bYQy7YyFQ
zq/mSBsuVv2XL2pkir4URY6W8/WTixCwwGT1xjomVX61da7ZR2bhs1zh+gTCE6Cm
0zPUesFd6r9gmZF3SjFw6xvOqoFo3n5Jlq138JDV1H7KY6xKtALTZH6A7qJgvVUn
CdL28tuBAgMBAAGjVDBSMA4GA1UdDwEB/wQEAwIApDATBgNVHSUEDDAKBggrBgEF
BQcDATAPBgNVHRMBAf8EBTADAQH/MBoGA1UdEQQTMBGCCWxvY2FsaG9zdIcEfwAA
ATALBgkqhkiG9w0BAQUDggEBAELOpAaNQy24GqwhJ4dXipsfIRUC8XmG0fn1Rc83
j6QHSsq/uBuC0492ScZ7rd4KLLKvJdnQd++2eGWsIvmGYy5sUS+X1Tm1twZ1ff8L
fiPFmLn+CymercgstHSo1fTXztunJmIs49uNvQ5E5Wjfb9bd9UK2pbQpoGRoDZrX
JKnW7YKqIkfHKnpgz5DtNvzn+dodvFvSV2TxxNyvHZYiJ93i4+/6GIqhAfEpkuym
3XQxGFYexX/ndHdEr5s3FTKlR20XB4sNWUyMmNIGYGQJD6SA8MNlYUH4VpDEIm88
JtuSOO2LKgSs2/emfjWHML5Zs3sz/YqOr1qJ2PJQTmd3Z4A=
-----END CERTIFICATE-----`