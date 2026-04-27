package armappnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appnetwork/armappnetwork"
)

// Generated from example definition: 2025-08-01-preview/AppLinkMembers_CreateOrUpdate.json
func ExampleAppLinkMembersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappnetwork.NewClientFactory("11809CA1-E126-4017-945E-AA795CD5C5A9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppLinkMembersClient().BeginCreateOrUpdate(ctx, "test_rg", "applink-test-01", "member-01", armappnetwork.AppLinkMember{
		Properties: &armappnetwork.AppLinkMemberProperties{
			ClusterType: to.Ptr(armappnetwork.ClusterTypeAKS),
			Metadata: &armappnetwork.Metadata{
				ResourceID: to.Ptr("/subscriptions/bc7e0da9-5e4c-4a91-9252-9658837006cf/resourcegroups/applink-rg/providers/Microsoft.ContainerService/managedClusters/applink-member1"),
			},
			UpgradeProfile: &armappnetwork.UpgradeProfile{
				Mode: to.Ptr(armappnetwork.UpgradeModeFullyManaged),
				FullyManagedUpgradeProfile: &armappnetwork.FullyManagedUpgradeProfile{
					ReleaseChannel: to.Ptr(armappnetwork.UpgradeReleaseChannelStable),
				},
			},
			ConnectivityProfile: &armappnetwork.ConnectivityProfile{
				PrivateConnect: &armappnetwork.PrivateConnectProfile{
					SubnetResourceID: to.Ptr("/subscriptions/bc7e0da9-5e4c-4a91-9252-9658837006cf/resourceGroups/applink-vnet-rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
				},
				EastWestGateway: &armappnetwork.EastWestGatewayProfile{
					Visibility: to.Ptr(armappnetwork.EastWestGatewayVisibilityInternal),
				},
			},
		},
		Tags: map[string]*string{
			"key2913": to.Ptr("test_tag"),
		},
		Location: to.Ptr("westus2"),
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
	// res = armappnetwork.AppLinkMembersClientCreateOrUpdateResponse{
	// 	AppLinkMember: &armappnetwork.AppLinkMember{
	// 		Properties: &armappnetwork.AppLinkMemberProperties{
	// 			ClusterType: to.Ptr(armappnetwork.ClusterTypeAKS),
	// 			Metadata: &armappnetwork.Metadata{
	// 				ResourceID: to.Ptr("/subscriptions/bc7e0da9-5e4c-4a91-9252-9658837006cf/resourcegroups/applink-rg/providers/Microsoft.ContainerService/managedClusters/applink-member1"),
	// 			},
	// 			UpgradeProfile: &armappnetwork.UpgradeProfile{
	// 				Mode: to.Ptr(armappnetwork.UpgradeModeFullyManaged),
	// 				FullyManagedUpgradeProfile: &armappnetwork.FullyManagedUpgradeProfile{
	// 					ReleaseChannel: to.Ptr(armappnetwork.UpgradeReleaseChannelStable),
	// 				},
	// 			},
	// 			ObservabilityProfile: &armappnetwork.ObservabilityProfile{
	// 				Metrics: &armappnetwork.MetricsProfile{
	// 					MetricsEndpoint: to.Ptr("https://mcp-fqdn"),
	// 				},
	// 			},
	// 			ConnectivityProfile: &armappnetwork.ConnectivityProfile{
	// 				PrivateConnect: &armappnetwork.PrivateConnectProfile{
	// 					SubnetResourceID: to.Ptr("/subscriptions/bc7e0da9-5e4c-4a91-9252-9658837006cf/resourceGroups/applink-vnet-rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 				},
	// 				EastWestGateway: &armappnetwork.EastWestGatewayProfile{
	// 					Visibility: to.Ptr(armappnetwork.EastWestGatewayVisibilityInternal),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armappnetwork.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key2913": to.Ptr("test_tag"),
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 		ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/test_rg/providers/Microsoft.AppLink/appLinks/applink-test-01/appLinkMembers/member-01"),
	// 		Name: to.Ptr("member-01"),
	// 		Type: to.Ptr("Microsoft.AppLink/appLinks/appLinkMembers"),
	// 		SystemData: &armappnetwork.SystemData{
	// 			CreatedBy: to.Ptr("user01"),
	// 			CreatedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user02"),
	// 			LastModifiedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
	// 		},
	// 	},
	// }
}
