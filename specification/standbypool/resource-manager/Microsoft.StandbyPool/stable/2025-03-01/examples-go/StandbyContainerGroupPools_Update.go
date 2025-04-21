package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool/v2"
)

// Generated from example definition: 2025-03-01/StandbyContainerGroupPools_Update.json
func ExampleStandbyContainerGroupPoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyContainerGroupPoolsClient().Update(ctx, "rgstandbypool", "pool", armstandbypool.StandbyContainerGroupPoolResourceUpdate{
		Tags: map[string]*string{},
		Properties: &armstandbypool.StandbyContainerGroupPoolResourceUpdateProperties{
			ElasticityProfile: &armstandbypool.StandbyContainerGroupPoolElasticityProfile{
				MaxReadyCapacity: to.Ptr[int64](1743),
				RefillPolicy:     to.Ptr(armstandbypool.RefillPolicyAlways),
			},
			ContainerGroupProperties: &armstandbypool.ContainerGroupProperties{
				ContainerGroupProfile: &armstandbypool.ContainerGroupProfile{
					ID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile"),
					Revision: to.Ptr[int64](2),
				},
				SubnetIDs: []*armstandbypool.Subnet{
					{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"),
					},
				},
			},
			Zones: []*string{
				to.Ptr("1"),
				to.Ptr("2"),
				to.Ptr("3"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstandbypool.StandbyContainerGroupPoolsClientUpdateResponse{
	// 	StandbyContainerGroupPoolResource: &armstandbypool.StandbyContainerGroupPoolResource{
	// 		Properties: &armstandbypool.StandbyContainerGroupPoolResourceProperties{
	// 			ElasticityProfile: &armstandbypool.StandbyContainerGroupPoolElasticityProfile{
	// 				MaxReadyCapacity: to.Ptr[int64](1743),
	// 				RefillPolicy: to.Ptr(armstandbypool.RefillPolicyAlways),
	// 			},
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 			ContainerGroupProperties: &armstandbypool.ContainerGroupProperties{
	// 				ContainerGroupProfile: &armstandbypool.ContainerGroupProfile{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.ContainerInstance/containerGroupProfiles/cgProfile"),
	// 					Revision: to.Ptr[int64](2),
	// 				},
	// 				SubnetIDs: []*armstandbypool.Subnet{
	// 					{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Network/virtualNetworks/cgSubnet/subnets/cgSubnet"),
	// 					},
	// 				},
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1"),
	// 				to.Ptr("2"),
	// 				to.Ptr("3"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyContainerGroupPools/pool"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		},
	// 	},
	// }
}
