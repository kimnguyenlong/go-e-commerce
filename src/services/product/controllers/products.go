package controllers

import (
	"context"
	"ecommerce/product/customerror"
	"ecommerce/product/entities"
	"ecommerce/product/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductsController struct {
	ProductModel *models.Product
}

func NewProductsController() (*ProductsController, error) {
	productModel, err := models.NewProduct()
	if err != nil {
		return nil, err
	}
	return &ProductsController{
		ProductModel: productModel,
	}, nil
}

func (pC ProductsController) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providerId := ctx.GetString("uid")
		var product entities.Product
		err := ctx.BindJSON(&product)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Bad request"))
		}
		product.ProviderID = providerId
		result, err := pC.ProductModel.Collection.InsertOne(context.Background(), product)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		product.ID = result.InsertedID.(primitive.ObjectID).Hex()
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  product,
		})
	}
}

func (pC ProductsController) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []*entities.Product
		cursor, err := pC.ProductModel.Collection.Find(context.Background(), bson.M{})
		if err != nil {
			panic(err)
		}
		defer cursor.Close(context.Background())
		err = cursor.All(context.Background(), &products)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  products,
		})

		// go to cache the products
		ctx.Set("products", products)
		ctx.Next()
	}
}

func (pC ProductsController) GetSingleProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		var product entities.Product
		filter := bson.M{"_id": pid}
		err = pC.ProductModel.Collection.FindOne(context.Background(), filter).Decode(&product)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No product with id: %s", pid.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  product,
		})
	}
}

func (pC ProductsController) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providerId := ctx.GetString("uid")
		productId, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		var product entities.Product
		filter := bson.M{
			"_id":        productId,
			"providerId": providerId,
		}
		err = pC.ProductModel.Collection.FindOneAndDelete(context.Background(), filter).Decode(&product)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No product with id: %s", productId.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  product,
		})
	}
}

func (pC ProductsController) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providerId := ctx.GetString("uid")
		productId, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		var updateData struct {
			Title       string   `json:"title" bson:"title,omitempty"`
			Price       float64  `json:"price" bson:"price,omitempty"`
			Categories  []string `json:"categories" bson:"categories,omitempty"`
			Description string   `json:"description" bson:"description,omitempty"`
		}
		err = ctx.BindJSON(&updateData)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}

		filter := bson.M{
			"_id":        productId,
			"providerId": providerId,
		}

		updateDataBytes, err := bson.Marshal(updateData)
		if err != nil {
			panic(err)
		}
		var bsonUpdate bson.M
		err = bson.Unmarshal(updateDataBytes, &bsonUpdate)
		if err != nil {
			panic(err)
		}
		update := bson.M{"$set": bsonUpdate}

		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		var product entities.Product
		err = pC.ProductModel.Collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&product)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No product with id: %s", productId.Hex())))
			}
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  product,
		})
	}
}
