// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package credfilecrypt

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var uname = (credfileLen() % 2) > 0
var pword = (credfileLen() % 2) == 0

type Credentials struct {
	f             *os.File
	CredFormat    *Credential
	sealed        bool
	enclaveBuffer *bytes.Buffer
}

type EncryptedCredentials struct {
	data   Crypter
	pwhash []byte // hash of bcrypt password
}

type Crypter interface {
	// CiphertextState should return the current encrypted state.
	// implementations should return true if in ciphertext (encrypted),
	// or false if in cleartext (unencrypted) state.
	CiphertextState() bool

	// Encrypt encrypts using pw.
	// If pw is nil, Encrypt will reuse the pw used last time it was called.
	// It returns non-nil error if pw == nil and this is the first invocation.
	Encrypt(pw []byte) error

	// Decrypt decrypts the ciphertext using pw.
	// Returns a nil error only when decryption is compleltely sucessful.
	Decrypt(pw []byte) error
}

type Credential struct {
	// where account exists. ex "google" for mail.google.com
	domain string `json:"host_domain"`

	// ex: "example@gmail.com"
	accountEmail string `json:"email"`

	// The password string generated by passgen utility
	generatedPass string `json:"passgen_string"`
}

func OpenCredentialsFile(file string) (*Credentials, error) {
	var f *Credentials
	var err error
	f.f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		// TODO: refactor this out into its own function
		if os.IsExist(err) {
			if credfilepath != "" {
				if !*createfileforce {
					return nil, fmt.Errorf("cannot create file %s, it already exists! (pass --force to force overwrite)", file)
				}
				if err := os.Remove(file); err != nil {

				}
				nf, _ := os.Create(file)
				// have to do this because this lays inside 'return nil, err'
				f = nf
				return f, nil
			}
		}
		return nil, err
	}
	p := &credfile{file: f}
	return p, nil
}

func (f *Credentials) Close() error {
	return f.Close()
	// TODO: refactor this out into its own function
}

func (f *Credentials) Decrypt(password []byte) error {
	bcrypt.CompareHashAndPassword()
	n, err := f.UnmarshalHex(f.enclaveBuffer)
	if err != nil {
		return err
	}

}

func (f *Credentials) CiphertextState() bool {
	return f.sealed
}

func (f *Credentials) unmarshalHex(b []byte) (n int, err error) {
	return io.ReadFull(f.f, b)

}