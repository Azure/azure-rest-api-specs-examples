package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a8698bb86e66e2d29ce5e8987b6aaa8fc7f7f04b/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-03-01-preview/examples/RedisEnterpriseListByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourceGroupPager("rg1", nil)
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
		// 				RedisVersion: to.Ptr("5"),
		// 				ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
		// 			},
		// 			SKU: &armredisenterprise.SKU{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 				Capacity: to.Ptr[int32](3),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1"),
		// 				to.Ptr("2"),
		// 				to.Ptr("3")},
		// 		}},
		// 	}
	}
}
