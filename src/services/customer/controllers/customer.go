package controllers

import (
	"context"
	"ecommerce/customer/customerror"
	"ecommerce/customer/entities"
	"ecommerce/customer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type CustomerController struct {
	CustomerModel *models.Customer
}

func NewCustomerController(dbCon *mongo.Database) (*CustomerController, error) {
	customerModel, err := models.NewCustomer(dbCon)
	if err != nil {
		return nil, err
	}
	return &CustomerController{
		CustomerModel: customerModel,
	}, nil
}

func (this CustomerController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var customer entities.Customer
		err := ctx.BindJSON(&customer)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Bad request"))
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		customer.Password = string(hashedPassword)
		result, err := this.CustomerModel.Collection.InsertOne(context.Background(), customer)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		customer.ID = result.InsertedID.(primitive.ObjectID).Hex()
		jwtString, _ := customer.CreateJWT()
		ctx.JSON(http.StatusCreated, gin.H{
			"error":     false,
			"data":      customer,
			"jwtString": jwtString,
		})
	}
}

func (this CustomerController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials struct {
			Email    string `json:"email" bson:"email"`
			Password string `json:"password" bson:"password"`
		}
		err := ctx.BindJSON(&credentials)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Bad request"))
		}
		var customer entities.Customer
		err = this.CustomerModel.Collection.FindOne(context.Background(), bson.M{"email": credentials.Email}).Decode(&customer)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusBadRequest, fmt.Sprintf("No Customer with email: %s", credentials.Email)))

			}
			panic(err)
		}
		if !customer.CheckPassword(credentials.Password) {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Incorrect password"))
		}
		jwtString, err := customer.CreateJWT()
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error":     false,
			"jwtString": jwtString,
		})
	}
}
