package security

import (
	"testing"
)

func TestHashPassword(t *testing.T) {

	password := "testpassword"

	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Fatalf("error hashing password: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatalf("hashed password is empty")
	}

	err = CheckPassword(string(hashedPassword), password)

	if err != nil {
		t.Fatalf("Password verification failed: %v", err)
	}

	err = CheckPassword(string(hashedPassword), "wrong password")

	if err == nil {
		t.Fatalf("Password verification should have failed but it did not for incorrect password")
	}

}
