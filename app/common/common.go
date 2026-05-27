package common

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

var MaxPassAttempts = 3

type PassVerifier func(passw []byte) (bool, error)

func VerifyPass(prompt string, verifier PassVerifier) error {
	for i := 0; i < MaxPassAttempts; i++ {
		fmt.Print(prompt)

		if p, err := term.ReadPassword(int(os.Stdin.Fd())); err != nil {
			return nil
		} else if ok, err := verifier(p); err != nil || ok {
			return err
		}
	}
	return errors.New("incorrect password attempts")
}
