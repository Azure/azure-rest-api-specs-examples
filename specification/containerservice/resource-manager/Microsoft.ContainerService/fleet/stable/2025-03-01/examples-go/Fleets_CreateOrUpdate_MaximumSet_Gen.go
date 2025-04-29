package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/Fleets_CreateOrUpdate_MaximumSet_Gen.json
func ExampleFleetsClient_BeginCreate_createsAFleetResourceWithALongRunningOperationGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginCreate(ctx, "rgfleets", "fleet1", armcontainerservicefleet.Fleet{
		Tags:     map[string]*string{},
		Location: to.Ptr("East US"),
		Properties: &armcontainerservicefleet.FleetProperties{
			HubProfile: &armcontainerservicefleet.FleetHubProfile{
				DNSPrefix: to.Ptr("dnsprefix1"),
				AgentProfile: &armcontainerservicefleet.AgentProfile{
					VMSize:   to.Ptr("Standard_DS1"),
					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
				},
				APIServerAccessProfile: &armcontainerservicefleet.APIServerAccessProfile{
					EnablePrivateCluster:  to.Ptr(true),
					EnableVnetIntegration: to.Ptr(true),
					SubnetID:              to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
				},
			},
		},
		Identity: &armcontainerservicefleet.ManagedServiceIdentity{
			Type: to.Ptr(armcontainerservicefleet.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armcontainerservicefleet.UserAssignedIdentity{
				"key126": {},
			},
		},
	}, &armcontainerservicefleet.FleetsClientBeginCreateOptions{
		IfMatch:     to.Ptr("jzlrwaylijhsnzp"),
		IfNoneMatch: to.Ptr("cqpzdjshmggwolagomzxfy")})
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
	// res = armcontainerservicefleet.FleetsClientCreateResponse{
	// 	Fleet: &armcontainerservicefleet.Fleet{
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armcontainerservicefleet.FleetProperties{
	// 			HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 				DNSPrefix: to.Ptr("dnsprefix1"),
	// 				AgentProfile: &armcontainerservicefleet.AgentProfile{
	// 					VMSize: to.Ptr("Standard_DS1"),
	// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
	// 				},
	// 				APIServerAccessProfile: &armcontainerservicefleet.APIServerAccessProfile{
	// 					EnablePrivateCluster: to.Ptr(true),
	// 					EnableVnetIntegration: to.Ptr(true),
	// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.Network/virtualNetwork/myvnet/subnets/mysubnet1"),
	// 				},
	// 				Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 				KubernetesVersion: to.Ptr("1.22.4"),
	// 				PortalFqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
	// 		},
	// 		ETag: to.Ptr("23ujdflewrj3="),
	// 		Identity: &armcontainerservicefleet.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Type: to.Ptr(armcontainerservicefleet.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armcontainerservicefleet.UserAssignedIdentity{
	// 				"key126": &armcontainerservicefleet.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet1"),
	// 		Name: to.Ptr("fleet-1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		},
	// 	},
	// }
}
