package ui_settings

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


// dummy data
var entryName1 = "Alex"
var entryName2 = "Bob"
var merchantCode = "Alex123"
var userId = uint64(1234567890)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017/ui-settings-test"))

	if err != nil {
		fmt.Println("error connecting MongoDB", err)
	}

	log.Println("mongoDB connected!")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("MongoDB ping failed", err)
	}

	log.Println("mongoDB pinged successfully!")
	return mongoTestClient
}

func cleanDatabase(t *testing.T, uiSettingRepo UiSettingRepo) func() {
	return func() {
		uiSettingRepo.DeleteEntryByName(merchantCode, entryName1)
		uiSettingRepo.DeleteEntryByName(merchantCode, entryName2)
		t.Log("Database cleared 🤌")
	}
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// connect to collection
	coll := mongoTestClient.Database("ui-settings-test").Collection("entries_test")

	uiSettingsRepo := UiSettingRepo{MongoCollection: coll}

	cleanup := cleanDatabase(t, uiSettingsRepo)
	defer cleanup()

	// insert entry #1
	t.Run("Insert entry #1", func (t *testing.T) {
		entry := UiSettingEntry{
			Name: entryName1,
			Value: 123,
			MerchantCode: merchantCode,
		}

		result, err := uiSettingsRepo.InsertEntry(&entry)
		if err != nil {
			t.Fatal("Insert entry #1 operation failed 🤯", err)
		}

		t.Log("Insert entry #1 💪", result)
	})

	// insert entry #2
	t.Run("Insert entry #2", func (t *testing.T) {
		entry := UiSettingEntry{
			Name: entryName2,
			Value: 123,
			MerchantCode: merchantCode,
			UserId: userId,
		}

		result, err := uiSettingsRepo.InsertEntry(&entry)
		if err != nil {
			t.Fatal("Insert entry #2 operation failed 🤯", err)
		}

		t.Log("Insert entry #2 💪", result)
	})

	// Get by merchant code
	t.Run("Get by merchant code", func (t *testing.T) {
		result, err := uiSettingsRepo.FindEntriesByMerchantCode(merchantCode)
		if err != nil {
			t.Fatal("Get by merchant code operation failed 🤯", err)
		}

		t.Log("Get by merchant code 💪", result)
		assert.Len(t, result, 2)
		assert.Equal(t, result[0].Name, entryName1)
	})

	// Get by user id
	t.Run("Get by user id", func (t *testing.T) {
		result, err := uiSettingsRepo.FindEntriesByUserId(merchantCode, userId)
		if err != nil {
			t.Fatal("Get by user id operation failed 🤯", err)
		}

		t.Log("Get by user id 💪", result)
		assert.Len(t, result, 1)
		assert.Equal(t, result[0].Name, entryName2)
	})
}