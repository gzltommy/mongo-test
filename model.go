package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"strconv"
	"time"
)

type mgo struct {
	collection string
}

func NewMgo(collection string) *mgo {
	return &mgo{
		collection,
	}
}

// 查询单个
func (m *mgo) FindOne(key string, value interface{}) *mongo.SingleResult {
	collection := db.Collection(m.collection).Clone()
	//collection.
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult
}

// 插入单个
func (m *mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	collection := db.Collection(m.collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		fmt.Println(err)
	}
	return insertResult
}

// 查询集合里有多少数据
func (m *mgo) CollectionCount() (string, int64) {
	collection := db.Collection(m.collection)
	name := collection.Name()
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

// 按选项查询集合 Skip 跳过 Limit 读取数量 sort 1 ，-1 . 1 为最初时间读取 ， -1 为最新时间读取
func (m *mgo) CollectionDocuments(Skip, Limit int64, sort int) *mongo.Cursor {
	collection := db.Collection(m.collection)
	SORT := bson.D{{"_id", sort}} //filter := bson.D{{key,value}}
	filter := bson.D{{}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	//findOptions.SetLimit(i)
	temp, _ := collection.Find(context.Background(), filter, findOptions)
	return temp
}

// 获取集合创建时间和编号
func (m *mgo) ParsingId(result string) (time.Time, uint64) {
	temp1 := result[:8]
	timestamp, _ := strconv.ParseInt(temp1, 16, 64)
	dateTime := time.Unix(timestamp, 0) //这是截获情报时间 时间格式 2019-04-24 09:23:39 +0800 CST
	temp2 := result[18:]
	count, _ := strconv.ParseUint(temp2, 16, 64) //截获情报的编号
	return dateTime, count
}

// 删除文章和查询文章
func (m *mgo) DeleteAndFind(key string, value interface{}) (int64, *mongo.SingleResult) {
	collection := db.Collection(m.collection)
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	DeleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println("删除时出现错误，你删不掉的~")
	}
	return DeleteResult.DeletedCount, singleResult
}

// 删除文章
func (m *mgo) Delete(key string, value interface{}) int64 {
	collection := db.Collection(m.collection)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount

}

// 删除多个
func (m *mgo) DeleteMany(key string, value interface{}) int64 {
	collection := db.Collection(m.collection)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount
}
