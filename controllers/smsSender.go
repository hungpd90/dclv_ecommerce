package controllers

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SmsSender() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response models.Response
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var messages models.MessageSend
		if err := c.BindJSON(&messages); err != nil {
			response.Status = "Failed"
			response.Code = http.StatusBadRequest
			response.Msg = err.Error()
			c.IndentedJSON(http.StatusBadRequest, response)
			return
		}
		accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
		authToken := os.Getenv("TWILIO_TOKEN")
		client := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})
		params := &twilioApi.CreateMessageParams{}
		params.SetTo(messages.PhoneNumberTo)
		params.SetFrom(os.Getenv("TWILIO_FROM_PHONE_NUMBER"))
		params.SetBody(messages.BodyMessage)
		_, anyErr := MessageCollection.InsertOne(ctx, messages)
		if anyErr != nil {
			response.Status = "Failed"
			response.Code = http.StatusInternalServerError
			response.Msg = "Not created"
			c.IndentedJSON(http.StatusInternalServerError, response)
			return
		}
		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			response, _ := json.Marshal(*resp)
			fmt.Println("Response: " + string(response))
		}

		response.Status = "OK"
		response.Code = http.StatusOK
		response.Msg = "Messages have been successfully sent by customer"
		c.IndentedJSON(http.StatusOK, response)
		return
	}
}
