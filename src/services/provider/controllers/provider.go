package controllers

import (
	"context"
	"ecommerce/provider/customerror"
	"ecommerce/provider/entities"
	"ecommerce/provider/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ProvidersController struct {
	ProviderModel *models.Provider
}

func NewProvidersController() (*ProvidersController, error) {
	providerModel, err := models.NewProvider()
	if err != nil {
		return nil, err
	}
	return &ProvidersController{
		ProviderModel: providerModel,
	}, nil
}

func (this ProvidersController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var provider entities.Provider
		err := ctx.BindJSON(&provider)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Bad request"))
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(provider.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		provider.Password = string(hashedPassword)
		result, err := this.ProviderModel.Collection.InsertOne(context.Background(), provider)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		provider.ID = result.InsertedID.(primitive.ObjectID).Hex()
		jwtString, _ := provider.CreateJWT()
		ctx.JSON(http.StatusCreated, gin.H{
			"error":     false,
			"data":      provider,
			"jwtString": jwtString,
		})
	}
}

func (this ProvidersController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials struct {
			Email    string `json:"email" bson:"email"`
			Password string `json:"password" bson:"password"`
		}
		err := ctx.BindJSON(&credentials)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Bad request"))
		}
		var provider entities.Provider
		err = this.ProviderModel.Collection.FindOne(context.Background(), bson.M{"email": credentials.Email}).Decode(&provider)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusBadRequest, fmt.Sprintf("No provider with email: %s", credentials.Email)))

			}
			panic(err)
		}
		if !provider.CheckPassword(credentials.Password) {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Incorrect password"))
		}
		jwtString, err := provider.CreateJWT()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error":     false,
			"jwtString": jwtString,
		})
	}
}
