package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad9b489baef1d982f7641f6c47a00794c9a1a5be/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armredisenterprise.OperationListResult{
		// 	Value: []*armredisenterprise.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/read"),
		// 			Display: &armredisenterprise.OperationDisplay{
		// 				Description: to.Ptr("View the Redis Enterprise cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache (read)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Cache/redisEnterprise/write"),
		// 			Display: &armredisenterprise.OperationDisplay{
		// 				Description: to.Ptr("Modify the Redis Enterprise cache's settings and configuration in the management portal"),
		// 				Operation: to.Ptr("Manage Redis Enterprise cache (write)"),
		// 				Provider: to.Ptr("Microsoft Cache"),
		// 				Resource: to.Ptr("Redis Enterprise cache"),
		// 			},
		// 	}},
		// }
	}
}
