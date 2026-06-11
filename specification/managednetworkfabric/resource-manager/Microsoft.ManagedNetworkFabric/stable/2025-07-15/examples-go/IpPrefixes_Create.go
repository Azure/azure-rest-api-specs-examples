package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/IpPrefixes_Create.json
func ExampleIPPrefixesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPPrefixesClient().BeginCreate(ctx, "example-rg", "example-ipPrefix", armmanagednetworkfabric.IPPrefix{
		Properties: &armmanagednetworkfabric.IPPrefixProperties{
			Annotation: to.Ptr("annotation"),
			IPPrefixRules: []*armmanagednetworkfabric.IPPrefixRule{
				{
					Action:           to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					SequenceNumber:   to.Ptr[int64](4155123341),
					NetworkPrefix:    to.Ptr("10.10.10.10/30"),
					Condition:        to.Ptr(armmanagednetworkfabric.ConditionGreaterThanOrEqualTo),
					SubnetMaskLength: to.Ptr("10"),
				},
			},
			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		},
		Tags: map[string]*string{
			"KeyId": to.Ptr("KeyValue"),
		},
		Location: to.Ptr("eastus"),
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
	// res = armmanagednetworkfabric.IPPrefixesClientCreateResponse{
	// 	IPPrefix: &armmanagednetworkfabric.IPPrefix{
	// 		Properties: &armmanagednetworkfabric.IPPrefixProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 			IPPrefixRules: []*armmanagednetworkfabric.IPPrefixRule{
	// 				{
	// 					Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 					NetworkPrefix: to.Ptr("10.10.10.10/30"),
	// 					Condition: to.Ptr(armmanagednetworkfabric.ConditionGreaterThanOrEqualTo),
	// 					SubnetMaskLength: to.Ptr("10"),
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"KeyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
	// 		Name: to.Ptr("example-ipPrefix"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/ipPrefixes"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
