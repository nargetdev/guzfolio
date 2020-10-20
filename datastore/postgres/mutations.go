package postgres

import (
	"guzfolio/model"

	"golang.org/x/crypto/bcrypt"
)

func (ds dataStore) CreateUser(input model.CreateUserInput) (*model.User, error) {
	// secure password
	bytePassword := []byte(input.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email: input.Email,
		Password: string(passwordHash),
		Name:  input.Name,
	}
	result := ds.db.Create(user)
	return user, result.Error
}

func (ds dataStore) CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error) {
	// get user
	user := model.User{}
	result := ds.db.Where("email = ?", input.UserEmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// get currency
	currency := model.Currency{}
	result = ds.db.Where("code = ?", input.FiatCurrencyCode).First(&currency)
	if result.Error != nil {
		return nil, result.Error
	}

	// create portfolio
	portfolio := &model.Portfolio{
		User:       	user,
		FiatCurrency:   currency,
		Name:           input.Name,
	}
	result = ds.db.Create(portfolio)
	return portfolio, result.Error
}

func (ds dataStore) CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error) {
	currency := &model.Currency{
		Code: input.Code,
		Name: input.Name,
		Type: input.Type,
	}
	result := ds.db.Create(currency)
	return currency, result.Error
}
