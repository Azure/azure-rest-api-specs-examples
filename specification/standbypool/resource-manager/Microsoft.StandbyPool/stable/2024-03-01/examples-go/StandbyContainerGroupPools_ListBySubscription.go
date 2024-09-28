package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: 2024-03-01/StandbyContainerGroupPools_ListBySubscription.json
func ExampleStandbyContainerGroupPoolsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandbyContainerGroupPoolsClient().NewListBySubscriptionPager(nil)
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
		// page = armstandbypool.StandbyContainerGroupPoolsClientListBySubscriptionResponse{
		// 	StandbyContainerGroupPoolResourceListResult: armstandbypool.StandbyContainerGroupPoolResourceListResult{
		// 		Value: []*armstandbypool.StandbyContainerGroupPoolResource{
		// 			{
		// 				Properties: &armstandbypool.StandbyContainerGroupPoolResourceProperties{
		// 					ElasticityProfile: &armstandbypool.StandbyContainerGroupPoolElasticityProfile{
		// 						MaxReadyCapacity: to.Ptr[int64](688),
		// 						RefillPolicy: to.Ptr(armstandbypool.RefillPolicyAlways),
		// 					},
		// 					ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
		// 					ContainerGroupProperties: &armstandbypool.ContainerGroupProperties{
		// 						ContainerGroupProfile: &armstandbypool.ContainerGroupProfile{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile"),
		// 							Revision: to.Ptr[int64](1),
		// 						},
		// 						SubnetIDs: []*armstandbypool.Subnet{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyContainerGroupPools/pool"),
		// 				Name: to.Ptr("pool"),
		// 				Type: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools"),
		// 				SystemData: &armstandbypool.SystemData{
		// 					CreatedBy: to.Ptr("pooluser@microsoft.com"),
		// 					CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
