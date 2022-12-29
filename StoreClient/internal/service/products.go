package service

import (
	"encoding/json"
	"fmt"
	"github.com/RomanVas30/storeClient/internal/request"
)

func ShowProducts(token string) error {
	resp, err := request.JsonRequest(
		request.Get,
		"http://localhost:8000/store/products/all",
		nil,
		token,
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

	fmt.Println(result)
	return nil
}
