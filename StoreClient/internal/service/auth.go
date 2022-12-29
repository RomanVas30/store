package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/RomanVas30/storeClient/internal/request"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func SignInUI() (username string, token string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		return
	}

	password := string(bytePassword)

	token, err = signIn(strings.TrimSpace(username), strings.TrimSpace(password))
	return
}

func signIn(username string, password string) (string, error) {
	resp, err := request.JsonRequest(
		request.Post,
		"http://localhost:8000/auth/sign-in",
		map[string]string{"username": username, "password": password},
		"",
	)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("the request was not executed: %s", result["message"])
	}

	return fmt.Sprintf("Bearer %s", result["token"]), nil
}

func SignUpUI() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Create username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Create password: ")
	byteNewPassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		return err
	}

	fmt.Print("Repeat password: ")
	byteRepeatPassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		return err
	}

	if string(byteNewPassword) != string(byteRepeatPassword) {
		return fmt.Errorf("passwords don't match")
	}

	password := string(byteNewPassword)

	return signUp(strings.TrimSpace(name), strings.TrimSpace(username), strings.TrimSpace(password))
}

func signUp(name string, username string, password string) error {
	resp, err := request.JsonRequest(
		request.Post,
		"http://localhost:8000/auth/sign-up",
		map[string]string{"name": name, "username": username, "password": password},
		"",
	)
	if err != nil {
		return err
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("the request was not executed: %s", result["message"])
	}

	return nil
}
