package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01/StoragePools_Update_MaximumSet_Gen.json
func ExampleStoragePoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStoragePoolsClient().BeginUpdate(ctx, "rgpurestorage", "storagePoolname", armpurestorageblock.StoragePoolUpdate{
		Identity: &armpurestorageblock.ManagedServiceIdentity{
			Type: to.Ptr(armpurestorageblock.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armpurestorageblock.UserAssignedIdentity{
				"key4211": {},
			},
		},
		Tags: map[string]*string{
			"key9065": to.Ptr("ebgmkwxqewe"),
		},
		Properties: &armpurestorageblock.StoragePoolUpdateProperties{
			ProvisionedBandwidthMbPerSec: to.Ptr[int64](23),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.StoragePoolsClientUpdateResponse{
	// 	StoragePool: &armpurestorageblock.StoragePool{
	// 		Properties: &armpurestorageblock.StoragePoolProperties{
	// 			StoragePoolInternalID: to.Ptr("zcvzukcmphctpzrebsgtcr"),
	// 			AvailabilityZone: to.Ptr("vknyl"),
	// 			VnetInjection: &armpurestorageblock.VnetInjection{
	// 				SubnetID: to.Ptr("tnlctolrxdvnkjiphlrdxq"),
	// 				VnetID: to.Ptr("zbumtytyqwewjcyckwqchiypshv"),
	// 			},
	// 			DataRetentionPeriod: to.Ptr[int64](23),
	// 			ProvisionedBandwidthMbPerSec: to.Ptr[int64](17),
	// 			ProvisionedIops: to.Ptr[int64](3),
	// 			Avs: &armpurestorageblock.AzureVmwareService{
	// 				AvsEnabled: to.Ptr(true),
	// 				ClusterResourceID: to.Ptr("zekrdsarbkwcbvpzhmuwoazogziwms"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpurestorageblock.ProvisioningStateSucceeded),
	// 			ReservationResourceID: to.Ptr("xiowoxnbtcotutcmmrofvgdi"),
	// 		},
	// 		Identity: &armpurestorageblock.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("550e8400-e29b-41d4-a716-446655440000"),
	// 			TenantID: to.Ptr("550e8400-e29b-41d4-a716-446655440000"),
	// 			Type: to.Ptr(armpurestorageblock.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armpurestorageblock.UserAssignedIdentity{
	// 				"key4211": &armpurestorageblock.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("550e8400-e29b-41d4-a716-446655440000"),
	// 					ClientID: to.Ptr("550e8400-e29b-41d4-a716-446655440000"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key7593": to.Ptr("vsyiygyurvwlfaezpuqu"),
	// 		},
	// 		Location: to.Ptr("lonlc"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("imnjvmnlsmalufopyw"),
	// 		Type: to.Ptr("hjztnpxxisrllusazxy"),
	// 		SystemData: &armpurestorageblock.SystemData{
	// 			CreatedBy: to.Ptr("ruoitchmuomrbscg"),
	// 			CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.341Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("thfyhokbrldzmghuylqbwpbublj"),
	// 			LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.345Z"); return t}()),
	// 		},
	// 	},
	// }
}
