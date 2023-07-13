package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	talon "github.com/talon-one/talon_go/v3"
)

func printStuff(v interface{}, response *http.Response) {
	if v != nil {
		fmt.Printf("%+v\n\n", v)
	}

	if response != nil {
		fmt.Printf("%+v\n\n", response)
	}
}

func client() *talon.APIClient {
	configuration := talon.NewConfiguration()
	// Set API base path
	configuration.Servers = talon.ServerConfigurations{
		{
			// Notice that there is no trailing '/'
			URL:         "http://localhost:9000",
			Description: "Talon.One's API base URL",
		},
	}

	client := talon.NewAPIClient(configuration)

	return client
}

func initIntegrationAPI() (*talon.IntegrationApiService, context.Context) {
	client := client()

	// Create integration authentication context using api key
	authContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"Authorization": {
			Prefix: "ApiKey-v1",
			Key:    os.Getenv("IAPI_KEY"),
		},
	})

	return client.IntegrationApi, authContext
}

func initManagementAPI() (*talon.ManagementApiService, context.Context) {
	client := client()
	// Create integration authentication context using api key
	authContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"Authorization": {
			Prefix: "ManagementKey-v1",
			Key:    os.Getenv("MAPI_KEY"),
		},
	})

	return client.ManagementApi, authContext
}

func TestGetApplication(t *testing.T) {
	api, context := initManagementAPI()

	application, response, err := api.
		GetApplication(context, 1).
		Execute()
	if err != nil {
		t.Fatal(fmt.Errorf("ERROR while calling GetApplication: %w", err))
	}

	printStuff(application, response)
}

func TestExportLoyalty(t *testing.T) {
	api, context := initManagementAPI()

	csvExport, response, err := api.
		ExportLoyaltyBalance(context, "1").
		Execute()
	if err != nil {
		panic(fmt.Errorf("ERROR while calling ExportLoyaltyBalance: %s\n", err))
	}

	printStuff(csvExport, response)
}

func TestManagementApi(t *testing.T) {
	t.Parallel()
	t.Run("TestExportLoyalty", func(t *testing.T) { TestExportLoyalty(t) })
	t.Run("TestGetApplication", func(t *testing.T) { TestGetApplication(t) })
}

func TestGetInventory(t *testing.T) {
	api, context := initIntegrationAPI()

	getInventoryRequest := api.GetCustomerInventory(context, "test_profile").
		Profile(true).
		Loyalty(true)

	inventory, _, err := getInventoryRequest.Execute()
	if err != nil {
		panic(fmt.Errorf("ERROR while calling GetCustomerInventory: %s\n", err))
	}

	printStuff(inventory.Profile, nil)
	printStuff(inventory.Loyalty, nil)
}
