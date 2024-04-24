package ui_settings

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UiSettingRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UiSettingRepo) InsertEntry(entry *UiSettingEntry) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), entry)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *UiSettingRepo) FindEntriesByMerchantCode(merchantCode string) ([]UiSettingEntry, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{{Key: "merchant_code", Value: merchantCode}})
	if err != nil {
		return nil, err
	}

	var entries []UiSettingEntry
	err = results.All(context.Background(), &entries)
	if err != nil {
		return nil, fmt.Errorf("decode error on FindEntriesByMerchantCode() %s", err)
	}

	return entries, nil
}

func (r *UiSettingRepo) FindEntriesByUserId(merchantCode string, userId uint64) ([]UiSettingEntry, error) {
	results, err := r.MongoCollection.Find(context.Background(),
		bson.D{
			{Key: "merchant_code", Value: merchantCode}, 
			{Key: "user_id", Value: userId},
		})
	if err != nil {
		return nil, err
	}

	var entries []UiSettingEntry
	err = results.All(context.Background(), &entries)
	if err != nil {
		return nil, fmt.Errorf("decode error on FindEntriesByUserId() %s", err)
	}

	return entries, nil
}

func (r *UiSettingRepo) UpdateEntryValueByName(merchantCode string, name string, value any) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{
			{Key: "name", Value: name},
			{Key: "merchant_code", Value: merchantCode},
		},
		bson.D{{Key: "$set", Value: updatedEntry}})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *UiSettingRepo) DeleteEntryByName(merchantCode string, name string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{
			{Key: "name", Value: name},
			{Key: "merchant_code", Value: merchantCode},
		})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}