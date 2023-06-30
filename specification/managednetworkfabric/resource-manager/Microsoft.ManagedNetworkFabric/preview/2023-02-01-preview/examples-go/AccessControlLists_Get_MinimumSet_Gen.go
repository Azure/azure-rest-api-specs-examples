package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/AccessControlLists_Get_MinimumSet_Gen.json
func ExampleAccessControlListsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessControlListsClient().Get(ctx, "resourceGroupName", "aclOne", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessControlList = armmanagednetworkfabric.AccessControlList{
	// 	Name: to.Ptr("aaaaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T07:15:42.793Z"); return t}()),
	// 		CreatedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-15T07:15:42.793Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("aaa"),
	// 	Properties: &armmanagednetworkfabric.AccessControlListProperties{
	// 		AddressFamily: to.Ptr(armmanagednetworkfabric.AddressFamilyIPv4),
	// 		Conditions: []*armmanagednetworkfabric.AccessControlListConditionProperties{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.ConditionActionTypeAllow),
	// 				DestinationAddress: to.Ptr("1.1.1.1"),
	// 				DestinationPort: to.Ptr("21"),
	// 				SequenceNumber: to.Ptr[int32](3),
	// 				SourceAddress: to.Ptr("2.2.2.2"),
	// 				SourcePort: to.Ptr("65000"),
	// 				Protocol: to.Ptr[int32](6),
	// 		}},
	// 	},
	// }
}
