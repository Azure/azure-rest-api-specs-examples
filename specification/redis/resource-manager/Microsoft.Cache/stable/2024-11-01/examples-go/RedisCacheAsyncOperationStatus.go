package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAsyncOperationStatus.json
func ExampleAsyncOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAsyncOperationStatusClient().Get(ctx, "East US", "c7ba2bf5-5939-4d79-b037-2964ccf097da", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armredis.OperationStatus{
	// 	Name: to.Ptr("c7ba2bf5-5939-4d79-b037-2964ccf097da"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Cache/locations/East US/asyncOperations/c7ba2bf5-5939-4d79-b037-2964ccf097da"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
