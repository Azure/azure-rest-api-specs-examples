package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_Update.json
func ExampleStorageAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakeanalytics.NewStorageAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"contosorg",
		"contosoadla",
		"test_storage",
		&armdatalakeanalytics.StorageAccountsClientUpdateOptions{Parameters: &armdatalakeanalytics.UpdateStorageAccountParameters{
			Properties: &armdatalakeanalytics.UpdateStorageAccountProperties{
				AccessKey: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
				Suffix:    to.Ptr("test_suffix"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
