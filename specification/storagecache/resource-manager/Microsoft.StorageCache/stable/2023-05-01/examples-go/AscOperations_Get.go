package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7f3e601fd326ca910c3d2939b516e15581e7e41/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/AscOperations_Get.json
func ExampleAscOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAscOperationsClient().Get(ctx, "westus", "testoperationid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AscOperation = armstoragecache.AscOperation{
	// 	Name: to.Ptr("testoperationid"),
	// 	EndTime: to.Ptr("2023-01-01T16:13:13.933Z"),
	// 	ID: to.Ptr("/subscriptions/id/locations/westus/ascOperations/testoperationid"),
	// 	StartTime: to.Ptr("2023-01-01T13:13:13.933Z"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
