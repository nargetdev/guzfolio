// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreatePortfolioInput struct {
	UserEmail        string  `json:"userEmail"`
	FiatCurrencyCode string  `json:"fiatCurrencyCode"`
	Name             *string `json:"name"`
}

type CreateUserInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}