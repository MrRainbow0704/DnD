package utils

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

// Creates a secret of given lenght.
func RandomSecret(l uint32) ([]byte, error) {
	s := make([]byte, l)
	_, err := rand.Read(s)
	if err != nil {
		return []byte{}, err
	}
	return s, nil
}

// Returns the hashed version of a password, given the unhashed version, salt and pepper.
func HashPassword(passwd, salt []byte) []byte {
	return argon2.IDKey(append(passwd, cnf.PasswdPepper...),
		salt,
		cnf.PasswdTime,
		cnf.PasswdMemory,
		cnf.PasswdThreads,
		cnf.PasswdKeyLen,
	)
}

// Reports weather the hashed password matches the clear password given the salt.
func ComparePassword(hash, passwd, salt []byte) bool {
	return bytes.Equal(hash, HashPassword(passwd, salt))
}
