package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PolicyInput struct {
	User     string `json:"user"`
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

type PolicyDecision struct {
	Result bool `json:"result"`
}

func queryOPA(user, action, resource string) (bool, error) {
	input := PolicyInput{
		User:     user,
		Action:   action,
		Resource: resource,
	}
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return false, err
	}

	fmt.Println(string(inputBytes))
	resp, err := http.Post("http://opa:8181/v1/data/authz/allow", "application/json", bytes.NewBuffer(inputBytes))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var decision PolicyDecision
	err = json.NewDecoder(resp.Body).Decode(&decision)
	if err != nil {
		return false, err
	}

	fmt.Println(decision)
	return decision.Result, nil
}

func main() {
	r := gin.Default()

	r.POST("/authorize", func(c *gin.Context) {
		var input PolicyInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		allowed, err := queryOPA(input.User, input.Action, input.Resource)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if allowed {
			c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		}
	})

	r.Run(":8080")
}
