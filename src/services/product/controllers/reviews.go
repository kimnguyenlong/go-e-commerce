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

type ReviewsController struct {
	reviewModel *models.Review
}

func NewReviewsController() (*ReviewsController, error) {
	reviewModel, err := models.NewReview()
	if err != nil {
		return nil, err
	}
	return &ReviewsController{
		reviewModel: reviewModel,
	}, nil
}

func (c ReviewsController) CreateReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		cid := ctx.GetString("uid")
		var review entities.Review
		err = ctx.BindJSON(&review)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		review.CustomerID = cid
		review.ProductID = pid.Hex()
		result, err := c.reviewModel.Collection.InsertOne(context.Background(), review)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		review.ID = result.InsertedID.(primitive.ObjectID).Hex()
		ctx.JSON(http.StatusCreated, gin.H{
			"error": false,
			"data":  review,
		})
	}
}

func (c ReviewsController) GetReviews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		filter := bson.M{"productId": pid.Hex()}
		curs, err := c.reviewModel.Collection.Find(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		defer curs.Close(context.Background())
		var reviews []*entities.Review
		err = curs.All(context.Background(), &reviews)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusOK, gin.H{
					"error": false,
					"data":  [...]entities.Review{},
				})
				return
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  reviews,
		})

		// go to cache the reviews
		ctx.Set("productId", pid.Hex())
		ctx.Set("reviews", reviews)
		ctx.Next()
	}
}

func (c ReviewsController) DeleteReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		rid, err := primitive.ObjectIDFromHex(ctx.Param("rid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid reviewId"))
		}
		cid := ctx.GetString("uid")
		filter := bson.M{
			"_id":        rid,
			"productId":  pid.Hex(),
			"customerId": cid,
		}
		var review *entities.Review
		err = c.reviewModel.Collection.FindOneAndDelete(context.Background(), filter).Decode(&review)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No review with id %s", rid.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  review,
		})
	}
}

func (c ReviewsController) UpdateReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
		}
		rid, err := primitive.ObjectIDFromHex(ctx.Param("rid"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid reviewId"))
		}
		cid := ctx.GetString("uid")
		var updateData struct {
			Content string `json:"content" bson:"content,omitempty" binding:"required"`
		}
		err = ctx.BindJSON(&updateData)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		filter := bson.M{
			"_id":        rid,
			"productId":  pid.Hex(),
			"customerId": cid,
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
		var review *entities.Review
		err = c.reviewModel.Collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&review)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No review with id %s", rid.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  review,
		})
	}
}
