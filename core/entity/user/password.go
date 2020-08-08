package user

// Implements Argon2ID password hashing and comparison. Heavily inspired on
// this article: http://www.inanzzz.com/index.php/post/0qug/hashing-and-verifying-passwords-with-argon2id-in-golang

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const hashFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
const saltLength = 16

type argon2ID struct {
	hashingTime uint32
	memory      uint32
	keyLength   uint32
	threads     uint8
}

// Hashes a plain text password using Argon2ID
func Hash(plainPassword string) (string, error) {
	argon := argon2ID{
		hashingTime: 1,
		memory:      64 * 1024,
		keyLength:   32,
		threads:     4,
	}

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(plainPassword), salt, argon.hashingTime, argon.memory, argon.threads, argon.keyLength)

	return fmt.Sprintf(
		hashFormat,
		argon2.Version,
		argon.memory,
		argon.hashingTime,
		argon.threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

// Verifies if a plain text password and a hash matches using Argon2ID
func Verify(plainPassword, hash string) (bool, error) {
	var argonVersion int
	var encodedSalt, encodedHash string
	argon := argon2ID{}

	_, err := fmt.Sscanf(
		hash,
		hashFormat,
		&argonVersion,
		&argon.memory,
		&argon.hashingTime,
		&argon.threads,
		&encodedSalt,
		&encodedHash,
	)

	if err != nil {
		return false, err
	}

	decodedSalt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(plainPassword), decodedSalt, argon.hashingTime, argon.memory, argon.threads, uint32(len(decodedHash)))
	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}
