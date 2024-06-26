package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpPrefixes_Create_MaximumSet_Gen.json
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
	poller, err := clientFactory.NewIPPrefixesClient().BeginCreate(ctx, "resourcegroupname", "example-ipPrefix", armmanagednetworkfabric.IPPrefix{
		Location: to.Ptr("EastUS"),
		Tags: map[string]*string{
			"key6404": to.Ptr(""),
		},
		Properties: &armmanagednetworkfabric.IPPrefixProperties{
			Annotation: to.Ptr("annotationValue"),
			IPPrefixRules: []*armmanagednetworkfabric.IPPrefixPropertiesIPPrefixRulesItem{
				{
					Action:           to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					Condition:        to.Ptr(armmanagednetworkfabric.ConditionEqualTo),
					NetworkPrefix:    to.Ptr("1.1.1.0/24"),
					SequenceNumber:   to.Ptr[int64](12),
					SubnetMaskLength: to.Ptr[int32](28),
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
	// 	ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T15:43:02.233Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-16T15:43:02.233Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("EastUS"),
	// 	Tags: map[string]*string{
	// 		"key6404": to.Ptr(""),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPPrefixProperties{
	// 		Annotation: to.Ptr("annotationValue"),
	// 		IPPrefixRules: []*armmanagednetworkfabric.IPPrefixPropertiesIPPrefixRulesItem{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				Condition: to.Ptr(armmanagednetworkfabric.ConditionEqualTo),
	// 				NetworkPrefix: to.Ptr("1.1.1.0/24"),
	// 				SequenceNumber: to.Ptr[int64](12),
	// 				SubnetMaskLength: to.Ptr[int32](28),
	// 		}},
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}
