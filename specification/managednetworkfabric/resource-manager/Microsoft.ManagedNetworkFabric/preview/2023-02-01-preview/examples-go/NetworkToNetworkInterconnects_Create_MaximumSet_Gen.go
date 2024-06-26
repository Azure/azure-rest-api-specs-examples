package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkToNetworkInterconnects_Create_MaximumSet_Gen.json
func ExampleNetworkToNetworkInterconnectsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkToNetworkInterconnectsClient().BeginCreate(ctx, "resourceGroupName", "FabricName", "DefaultNNI", armmanagednetworkfabric.NetworkToNetworkInterconnect{
		Properties: &armmanagednetworkfabric.NetworkToNetworkInterconnectProperties{
			IsManagementType: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
			Layer2Configuration: &armmanagednetworkfabric.Layer2Configuration{
				Mtu:       to.Ptr[int32](1500),
				PortCount: to.Ptr[int32](10),
			},
			Layer3Configuration: &armmanagednetworkfabric.Layer3Configuration{
				PrimaryIPv4Prefix:   to.Ptr("172.31.0.0/31"),
				PrimaryIPv6Prefix:   to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
				SecondaryIPv4Prefix: to.Ptr("172.31.0.20/31"),
				SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
				ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
				ImportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
				PeerASN:             to.Ptr[int32](50272),
				VlanID:              to.Ptr[int32](2064),
			},
			NniType:    to.Ptr(armmanagednetworkfabric.NniTypeCE),
			UseOptionB: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyFalse),
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
	// res.NetworkToNetworkInterconnect = armmanagednetworkfabric.NetworkToNetworkInterconnect{
	// 	Name: to.Ptr("DefaultNNI"),
	// 	Type: to.Ptr("Microsoft.ManagedNetworkFabric/networkFabrics/FabricName/networkToNetworkInterconnect/DefaultNNI"),
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName/networkToNetworkInterconnect/DefaultNNI"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-11T04:17:22.639Z"); return t}()),
	// 		CreatedBy: to.Ptr("User@email.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-11T04:17:22.639Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkToNetworkInterconnectProperties{
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
	// 		IsManagementType: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 		Layer2Configuration: &armmanagednetworkfabric.Layer2Configuration{
	// 			Interfaces: []*string{
	// 				to.Ptr("defaultInterface")},
	// 				Mtu: to.Ptr[int32](2516),
	// 				PortCount: to.Ptr[int32](3),
	// 			},
	// 			Layer3Configuration: &armmanagednetworkfabric.Layer3Configuration{
	// 				PrimaryIPv4Prefix: to.Ptr("172.31.0.0/31"),
	// 				PrimaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
	// 				SecondaryIPv4Prefix: to.Ptr("172.31.0.20/31"),
	// 				SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
	// 				ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2"),
	// 				FabricASN: to.Ptr[int32](1253),
	// 				ImportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1"),
	// 				PeerASN: to.Ptr[int32](50272),
	// 				VlanID: to.Ptr[int32](2064),
	// 			},
	// 			NniType: to.Ptr(armmanagednetworkfabric.NniTypeCE),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 			UseOptionB: to.Ptr(armmanagednetworkfabric.BooleanEnumPropertyTrue),
	// 		},
	// 	}
}
