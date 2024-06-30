package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/wire"
	"github.com/joho/godotenv"
)

var Set = wire.NewSet(NewConfig)

type Configuration struct {
	Server         ServerConfig
	MSSQL          MSSQLConfig
	MongoDB        MongoDBConfig
	Typesense      TypesenseConfig
	Google         GoogleConfig
	Elasticsearch  ElasticsearchConfig
	Connectors     ConnectorsConfig
	ContentManager ContentManagerConfig
	Encryption     EncryptionConfig
	ERP            ERPConfig
	Service            ServiceConfig
	ExchangeRate            ExchangeRateConfig
}
type ServerConfig struct {
	ENV                 string
	VERSION             string
	PORT                string
	HOST                string
	USERNAME            string
	SERVER_READ_TIMEOUT int
	STAGE_STATUS        string
}
type MSSQLConfig struct {
	IP             string
	USERNAME       string
	PASSWORD       string
	DB             string
	PORT           int
	MAX_OPEN_CONNS int
	MAX_IDLE_CONNS int
	CONN_TIME_OUT  int
	IDLE_TIME_OUT  int
	REQ_TIME_OUT   int
}
type MongoDBConfig struct {
	IP             string
	USERNAME       string
	PASSWORD       string
	DB             string
	PORT           int
	MAX_OPEN_CONNS int
	MAX_IDLE_CONNS int
	CONN_TIME_OUT  int
	IDLE_TIME_OUT  int
	REQ_TIME_OUT   int
}
type TypesenseConfig struct {
	API_KAY    string
	HOST       string
	PROTOCOL   string
	COLLECTION string
	PORT       int
}

type GoogleConfig struct {
	KEY string
}
type ElasticsearchConfig struct {
	ElasticUrl             string
	ElasticUsername        string
	ElasticPassword        string
	ElasticIndexDataSearch string
	ElasticTimeout  	     time.Duration
}

type ConnectorsConfig struct {
	Agoda     ConnectorConfig
	Dotw      ConnectorConfig
	Extranet  ConnectorConfig
	Hotelbeds ConnectorConfig
}
type ConnectorConfig struct {
	Endpoint    string
	SupplierKey string
	GetHotelAvailabilityTimeout  time.Duration
	GetRoomsAvailabilityTimeout  time.Duration
}
type ContentManagerConfig struct {
	Endpoint string
	GetRoomDetailTimeout  time.Duration
	GetHotelListTimeout  time.Duration
}
type EncryptionConfig struct {
	KeyBASE64 string
}
type ERPConfig struct {
	MarkPriceEndpoint string
	BuId             int
	ProductTypeId    int
	GetMarkingPriceTimeout     time.Duration
}
type ExchangeRateConfig struct {
	ExchangeRateEndpoint     string
	ExchangeRateAPIKey     string
	GetExchangeRateTimeout     time.Duration
}
type ServiceConfig struct {
	SearchHotelItemPerPage    int
	SearchHotelMaxItemPerPage    int
	BaseCurrency    string
}

func NewConfig() (*Configuration, error) {
	currentWorkDirectory, _ := os.Getwd()
	environment := "production"

	if os.Getenv("ENV") == "development" {
		environment = "development"
	}
	var envFilePath string = fmt.Sprintf("%s/env/%s.env", currentWorkDirectory, environment)
	fmt.Println("Environment:", environment)
	err := godotenv.Load(envFilePath)

	if err != nil {
		log.Fatal("[Error] loaded env file fail: ", err.Error())
	}
	config := GetConfiguration()
	return &config, nil
}

func GetConfiguration() Configuration {
	var ENV string = os.Getenv("ENV")
	var VERSION string = os.Getenv("VERSION")
	var HOST string = os.Getenv("HOST")
	var USERNAME string = os.Getenv("USERNAME")
	var PORT string = os.Getenv("PORT")
	var STAGE_STATUS string = os.Getenv("STAGE_STATUS")
	SERVER_READ_TIMEOUT, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	var MSSQL_IP string = os.Getenv("MSSQL_IP")
	var MSSQL_USERNAME string = os.Getenv("MSSQL_USERNAME")
	var MSSQL_PASSWORD string = os.Getenv("MSSQL_PASSWORD")
	var MSSQL_DB string = os.Getenv("MSSQL_DB")
	MSSQL_PORT, _ := strconv.Atoi(os.Getenv("MSSQL_PORT"))

	var MONGO_IP string = os.Getenv("MONGO_IP")
	var MONGO_USERNAME string = os.Getenv("MONGO_USERNAME")
	var MONGO_PASSWORD string = os.Getenv("MONGO_PASSWORD")
	var MONGO_DB string = os.Getenv("MONGO_DB")
	MONGO_PORT, _ := strconv.Atoi(os.Getenv("MONGO_PORT"))

	var TYPESENSE_API_KAY string = os.Getenv("TYPESENSE_API_KAY")
	var TYPESENSE_HOST string = os.Getenv("TYPESENSE_HOST")
	var TYPESENSE_PROTOCOL string = os.Getenv("TYPESENSE_PROTOCOL")
	var TYPESENSE_COLLECTION string = os.Getenv("TYPESENSE_COLLECTION")
	TYPESENSE_PORT, _ := strconv.Atoi(os.Getenv("TYPESENSE_PORT"))

	var ElasticUrl = os.Getenv("ELASTIC_URL")
	var ElasticUsername string = os.Getenv("ELASTIC_USERNAME")
	var ElasticPassword string = os.Getenv("ELASTIC_PASSWORD")
	var ElasticIndexDataSearch string = os.Getenv("ELASTIC_INDEX_DATA_SEARCH")

	var GOOGLE_KEY string = os.Getenv("GOOGLE_KEY")

	var CONTENT_SERVICE_ENDPOINT string = os.Getenv("CONTENT_SERVICE_ENDPOINT")

	var CONNECTOR_SUPPLIER_KEY string = os.Getenv("CONNECTOR_SUPPLIER_KEY")
	var CONNECTOR_HOTELBEDS_ENDPOINT string = os.Getenv("CONNECTOR_HOTELBEDS_ENDPOINT")
	var CONNECTOR_EXTRANET_ENDPOINT string = os.Getenv("CONNECTOR_EXTRANET_ENDPOINT")
	var CONNECTOR_AGODA_ENDPOINT string = os.Getenv("CONNECTOR_AGODA_ENDPOINT")
	var KeyBASE64 string = os.Getenv("ENCRYPTION_KEY_BASE64")
	var MarkPriceEndpoint string = os.Getenv("ERP_MARK_PRICE_ENDPOINT")

	ERP_BU_ID, _ := strconv.Atoi(os.Getenv("ERP_BU_ID"))
	ERP_PRODUCT_TYPE_ID, _ := strconv.Atoi(os.Getenv("ERP_PRODUCT_TYPE_ID"))

	EXCHANGE_RATE_ENDPOINT := os.Getenv("EXCHANGE_RATE_ENDPOINT")
	EXCHANGE_RATE_API_KEY := os.Getenv("EXCHANGE_RATE_API_KEY")
	// SEARCH_HOTEL_ITEM_PER_PAGE, _ := strconv.Atoi(os.Getenv("SEARCH_HOTEL_ITEM_PER_PAGE"))

	config := &Configuration{
		Server: ServerConfig{
			ENV:                 ENV,
			SERVER_READ_TIMEOUT: SERVER_READ_TIMEOUT,
			VERSION:             VERSION,
			PORT:                PORT,
			HOST:                HOST,
			USERNAME:            USERNAME,
			STAGE_STATUS:        STAGE_STATUS,
		},
		MSSQL: MSSQLConfig{
			IP:             MSSQL_IP,
			USERNAME:       MSSQL_USERNAME,
			PASSWORD:       MSSQL_PASSWORD,
			DB:             MSSQL_DB,
			PORT:           MSSQL_PORT,
			MAX_OPEN_CONNS: 5,
			MAX_IDLE_CONNS: 5,
			CONN_TIME_OUT:  32000,
			IDLE_TIME_OUT:  32000,
			REQ_TIME_OUT:   32000,
		},
		MongoDB: MongoDBConfig{
			IP:             MONGO_IP,
			USERNAME:       MONGO_USERNAME,
			PASSWORD:       MONGO_PASSWORD,
			DB:             MONGO_DB,
			PORT:           MONGO_PORT,
			MAX_OPEN_CONNS: 5,
			MAX_IDLE_CONNS: 5,
			CONN_TIME_OUT:  32000,
			IDLE_TIME_OUT:  32000,
			REQ_TIME_OUT:   32000,
		},
		Typesense: TypesenseConfig{
			API_KAY:    TYPESENSE_API_KAY,
			HOST:       TYPESENSE_HOST,
			PROTOCOL:   TYPESENSE_PROTOCOL,
			COLLECTION: TYPESENSE_COLLECTION,
			PORT:       TYPESENSE_PORT,
		},
		Google: GoogleConfig{
			KEY: GOOGLE_KEY,
		},
		Elasticsearch: ElasticsearchConfig{
			ElasticUrl: ElasticUrl,
			ElasticUsername: ElasticUsername,
			ElasticPassword: ElasticPassword,
			ElasticIndexDataSearch: ElasticIndexDataSearch,
			ElasticTimeout: time.Second * 5, // 5 seconds timeout
		},
		Connectors: ConnectorsConfig{
			Hotelbeds: ConnectorConfig{
				Endpoint:    CONNECTOR_HOTELBEDS_ENDPOINT,
				SupplierKey: CONNECTOR_SUPPLIER_KEY,
				GetHotelAvailabilityTimeout: time.Second * 3, // 3 seconds timeout
				GetRoomsAvailabilityTimeout: time.Second * 15, // 15 seconds timeout
			},
			Extranet: ConnectorConfig{
				Endpoint:    CONNECTOR_EXTRANET_ENDPOINT,
				SupplierKey: CONNECTOR_SUPPLIER_KEY,
				GetHotelAvailabilityTimeout: time.Second * 3, // 3 seconds timeout
				GetRoomsAvailabilityTimeout: time.Second * 15, // 15 seconds timeout
			},
			Agoda: ConnectorConfig{
				Endpoint:    CONNECTOR_AGODA_ENDPOINT,
				SupplierKey: CONNECTOR_SUPPLIER_KEY,
				GetHotelAvailabilityTimeout: time.Second * 3, // 3 seconds timeout
				GetRoomsAvailabilityTimeout: time.Second * 15, // 15 seconds timeout
			},
			Dotw: ConnectorConfig{},
		},
		ContentManager: ContentManagerConfig{
			Endpoint: CONTENT_SERVICE_ENDPOINT,
			GetRoomDetailTimeout:  time.Second * 3, // 3 seconds timeout,
			GetHotelListTimeout:  time.Second * 3, // 3 seconds timeout,
		},
		Encryption: EncryptionConfig{
			KeyBASE64: KeyBASE64,
		},
		ERP: ERPConfig{
			MarkPriceEndpoint: MarkPriceEndpoint,
			BuId:             ERP_BU_ID,
			ProductTypeId:    ERP_PRODUCT_TYPE_ID,
			GetMarkingPriceTimeout:  time.Second * 30 ,
		},
		ExchangeRate: ExchangeRateConfig{
			ExchangeRateEndpoint: EXCHANGE_RATE_ENDPOINT, 
			ExchangeRateAPIKey : EXCHANGE_RATE_API_KEY, 
			GetExchangeRateTimeout: time.Second * 3, // 3 seconds timeout
		},
		Service: ServiceConfig{
			SearchHotelItemPerPage: 40,
			SearchHotelMaxItemPerPage: 400,
			BaseCurrency: "THB",
		},
	}
	return *config
}
