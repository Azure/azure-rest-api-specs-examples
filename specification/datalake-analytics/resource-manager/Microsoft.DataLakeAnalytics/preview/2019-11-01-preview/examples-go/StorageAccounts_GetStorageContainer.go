package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_GetStorageContainer.json
func ExampleStorageAccountsClient_GetStorageContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageAccountsClient().GetStorageContainer(ctx, "contosorg", "contosoadla", "test_storage", "test_container", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageContainer = armdatalakeanalytics.StorageContainer{
	// 	Name: to.Ptr("test_storage"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakeanalytics.StorageContainerProperties{
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
	// 	},
	// }
}
