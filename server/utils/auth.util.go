package utils

import (
	"bytes"
	"crypto/rand"
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/argon2"
)

func RandomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func HashString(value string, salt []byte) ([]byte, error) {
	if len(value) == 0 {
		return nil, errors.New("could not hash string, given value len is 0")
	}
	if len(salt) == 0 {
		return nil, errors.New("could not hash string, salt len is 0")
	}

	time, err := strconv.ParseUint(os.Getenv("ARGON_TIME"), 10, 32)
	if err != nil {
		return nil, errors.New("failed to hash string, could not parse ARGON_TIME ENV variable")
	}
	memory, err := strconv.ParseUint(os.Getenv("ARGON_MEMORY"), 10, 32)
	if err != nil {
		return nil, errors.New("failed to hash string, could not parse ARGON_MEMORY ENV variable")
	}
	threads, err := strconv.ParseUint(os.Getenv("ARGON_THREADS"), 10, 8)
	if err != nil {
		return nil, errors.New("failed to hash string, could not parse ARGON_THREADS ENV variable")
	}
	keyLen, err := strconv.ParseUint(os.Getenv("AGON_KEY_LENGTH"), 10, 32)
	if err != nil {
		return nil, errors.New("failed to hash string, could not parse ARGON_KEY_LENGTH ENV variable")
	}

	hash := argon2.IDKey(
		[]byte(value),
		salt,
		uint32(time),
		uint32(memory),
		uint8(threads),
		uint32(keyLen),
	)

	return hash, nil
}

func CompareHash(h string, s string, v string) error {
	hashSalt, err := HashString(v, []byte(s))
	if err != nil {
		return err
	}

	if !bytes.Equal(hashSalt, []byte(h)) {
		return errors.New("hash does not match")
	}

	return nil
}
