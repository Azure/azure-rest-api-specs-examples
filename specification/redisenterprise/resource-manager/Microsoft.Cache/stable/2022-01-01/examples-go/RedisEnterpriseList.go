package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseList.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
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
		// page.ClusterList = armredisenterprise.ClusterList{
		// 	Value: []*armredisenterprise.Cluster{
		// 		{
		// 			Name: to.Ptr("cache1"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armredisenterprise.ClusterProperties{
		// 				HostName: to.Ptr("cache1.westus.something.azure.com"),
		// 				MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				RedisVersion: to.Ptr("6"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 			},
		// 			SKU: &armredisenterprise.SKU{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 				Capacity: to.Ptr[int32](3),
		// 			},
		// 	}},
		// }
	}
}
