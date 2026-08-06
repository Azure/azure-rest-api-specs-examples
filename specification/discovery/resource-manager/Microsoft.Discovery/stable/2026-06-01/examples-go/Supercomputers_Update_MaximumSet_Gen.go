package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Supercomputers_Update_MaximumSet_Gen.json
func ExampleSupercomputersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSupercomputersClient().BeginUpdate(ctx, "rgdiscovery", "1d951e48d0e7383455", armdiscovery.Supercomputer{
		Properties: &armdiscovery.SupercomputerProperties{
			Identities: &armdiscovery.SupercomputerIdentities{
				WorkloadIdentities: map[string]*armdiscovery.UserAssignedIdentity{
					"key7289": {},
				},
			},
		},
		Identity: &armdiscovery.SystemAssignedServiceIdentity{
			Type: to.Ptr(armdiscovery.SystemAssignedServiceIdentityTypeSystemAssigned),
		},
		Tags: map[string]*string{
			"key40": to.Ptr("guakqh"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.SupercomputersClientUpdateResponse{
	// 	Supercomputer: armdiscovery.Supercomputer{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/supercomputers/1d951e48d0e7383455"),
	// 		Name: to.Ptr("1d951e48d0e7383455"),
	// 		Tags: map[string]*string{
	// 			"key5117": to.Ptr("zeawptiwdzd"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/supercomputers"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Identity: &armdiscovery.SystemAssignedServiceIdentity{
	// 			PrincipalID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 			TenantID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 			Type: to.Ptr(armdiscovery.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		},
	// 		Properties: &armdiscovery.SupercomputerProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			SubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/subnet1"),
	// 			ManagementSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/managementSubnet1"),
	// 			OutboundType: to.Ptr(armdiscovery.NetworkEgressTypeLoadBalancer),
	// 			SystemSKU: to.Ptr(armdiscovery.SystemSKUStandardD4SV6),
	// 			Identities: &armdiscovery.SupercomputerIdentities{
	// 				ClusterIdentity: &armdiscovery.Identity{
	// 					ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedid1"),
	// 				},
	// 				KubeletIdentity: &armdiscovery.Identity{
	// 					ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedid1"),
	// 				},
	// 				WorkloadIdentities: map[string]*armdiscovery.UserAssignedIdentity{
	// 					"key8010": &armdiscovery.UserAssignedIdentity{
	// 						PrincipalID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 						ClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 					},
	// 				},
	// 			},
	// 			CustomerManagedKeys: to.Ptr(armdiscovery.CustomerManagedKeysEnabled),
	// 			DiskEncryptionSetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Compute/diskEncryptionSets/diskencryptionset1"),
	// 			LogAnalyticsClusterID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.OperationalInsights/clusters/cluster1"),
	// 			ManagedResourceGroup: to.Ptr("ltsdbtlcjzndpukxdmnykpguudw"),
	// 			ManagedOnBehalfOfConfiguration: &armdiscovery.WithMoboBrokerResources{
	// 				MoboBrokerResources: []*armdiscovery.MoboBrokerResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Storage/storageAccounts/storage1"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
