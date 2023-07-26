package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpPrefixes_Create_MaximumSet_Gen.json
func ExampleIPPrefixesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPPrefixesClient().BeginCreate(ctx, "example-rg", "example-ipPrefix", armmanagednetworkfabric.IPPrefix{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.IPPrefixProperties{
			Annotation: to.Ptr("annotation"),
			IPPrefixRules: []*armmanagednetworkfabric.IPPrefixRule{
				{
					Action:           to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					Condition:        to.Ptr(armmanagednetworkfabric.ConditionGreaterThanOrEqualTo),
					NetworkPrefix:    to.Ptr("10.10.10.10/30"),
					SequenceNumber:   to.Ptr[int64](4155123341),
					SubnetMaskLength: to.Ptr("10"),
				}},
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
	// res.IPPrefix = armmanagednetworkfabric.IPPrefix{
	// 	Name: to.Ptr("example-ipPrefix"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipPrefixes"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T03:52:05.656Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-12T03:52:05.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPPrefixProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		IPPrefixRules: []*armmanagednetworkfabric.IPPrefixRule{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				Condition: to.Ptr(armmanagednetworkfabric.ConditionGreaterThanOrEqualTo),
	// 				NetworkPrefix: to.Ptr("10.10.10.10/30"),
	// 				SequenceNumber: to.Ptr[int64](4155123341),
	// 				SubnetMaskLength: to.Ptr("10"),
	// 		}},
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}
