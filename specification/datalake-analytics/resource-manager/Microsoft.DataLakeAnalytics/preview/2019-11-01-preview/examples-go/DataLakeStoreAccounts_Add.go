package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/DataLakeStoreAccounts_Add.json
func ExampleDataLakeStoreAccountsClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewDataLakeStoreAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Add(ctx,
		"contosorg",
		"contosoadla",
		"test_adls_account",
		&armdatalakeanalytics.DataLakeStoreAccountsClientAddOptions{Parameters: &armdatalakeanalytics.AddDataLakeStoreParameters{
			Properties: &armdatalakeanalytics.AddDataLakeStoreProperties{
				Suffix: to.Ptr("test_suffix"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
