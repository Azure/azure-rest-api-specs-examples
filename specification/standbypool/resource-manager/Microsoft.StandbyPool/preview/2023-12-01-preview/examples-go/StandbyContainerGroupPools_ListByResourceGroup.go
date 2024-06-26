package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyContainerGroupPools_ListByResourceGroup.json
func ExampleStandbyContainerGroupPoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandbyContainerGroupPoolsClient().NewListByResourceGroupPager("rgstandbypool", nil)
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
		// page.StandbyContainerGroupPoolResourceListResult = armstandbypool.StandbyContainerGroupPoolResourceListResult{
		// 	Value: []*armstandbypool.StandbyContainerGroupPoolResource{
		// 		{
		// 			Name: to.Ptr("pool"),
		// 			Type: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools"),
		// 			ID: to.Ptr("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyContainerGroupPools/pool"),
		// 			SystemData: &armstandbypool.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				CreatedBy: to.Ptr("pooluser@microsoft.com"),
		// 				CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armstandbypool.StandbyContainerGroupPoolResourceProperties{
		// 				ContainerGroupProperties: &armstandbypool.ContainerGroupProperties{
		// 					ContainerGroupProfile: &armstandbypool.ContainerGroupProfile{
		// 						ID: to.Ptr("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile"),
		// 						Revision: to.Ptr[int64](1),
		// 					},
		// 					SubnetIDs: []*armstandbypool.Subnet{
		// 						{
		// 							ID: to.Ptr("/subscriptions/8cf6c1b6-c80f-437c-87ad-45fbaff54f73/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"),
		// 					}},
		// 				},
		// 				ElasticityProfile: &armstandbypool.StandbyContainerGroupPoolElasticityProfile{
		// 					MaxReadyCapacity: to.Ptr[int64](688),
		// 					RefillPolicy: to.Ptr(armstandbypool.RefillPolicyAlways),
		// 				},
		// 				ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
