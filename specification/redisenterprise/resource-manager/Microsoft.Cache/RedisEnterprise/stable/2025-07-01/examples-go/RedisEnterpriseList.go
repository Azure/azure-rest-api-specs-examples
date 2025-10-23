package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3855ffb4be0cd4d227b130b67d874fa816736c04/specification/redisenterprise/resource-manager/Microsoft.Cache/RedisEnterprise/stable/2025-07-01/examples/RedisEnterpriseList.json
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
		// 			ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Kind: to.Ptr(armredisenterprise.KindV1),
		// 			Properties: &armredisenterprise.ClusterCreateProperties{
		// 				HostName: to.Ptr("cache1.westus.something.azure.com"),
		// 				MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
		// 				ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
		// 				RedisVersion: to.Ptr("6"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 				PublicNetworkAccess: to.Ptr(armredisenterprise.PublicNetworkAccessDisabled),
		// 			},
		// 			SKU: &armredisenterprise.SKU{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 				Capacity: to.Ptr[int32](3),
		// 			},
		// 	}},
		// }
	}
}
