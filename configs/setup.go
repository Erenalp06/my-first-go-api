package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	mongoURI := EnvMongoURI()

	// MongoDB bağlantı ayarları
	clientOptions := options.Client().ApplyURI(mongoURI)

	// MongoDB istemcisini oluşturduk
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB'ye bağlanırken bir hata oluştu: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB sunucusna ping atılamadı: %v", err)
	}

	log.Println("MongoDB'ye başarıyla bağlanıldı.")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golandAPI").Collection(collectionName)
}

func SetupLogger() {
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}
