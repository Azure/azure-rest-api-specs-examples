package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_UpdateAdministrativeState_MaximumSet_Gen.json
func ExampleAccessControlListsClient_BeginUpdateAdministrativeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessControlListsClient().BeginUpdateAdministrativeState(ctx, "example-rg", "example-acl", armmanagednetworkfabric.UpdateAdministrativeState{
		ResourceIDs: []*string{
			to.Ptr("")},
		State: to.Ptr(armmanagednetworkfabric.EnableDisableStateEnable),
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
	// res.CommonPostActionResponseForStateUpdate = armmanagednetworkfabric.CommonPostActionResponseForStateUpdate{
	// 	Error: &armmanagednetworkfabric.ErrorDetail{
	// 		AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 			{
	// 				Info: map[string]any{
	// 				},
	// 				Type: to.Ptr(""),
	// 		}},
	// 		Code: to.Ptr(""),
	// 		Message: to.Ptr(""),
	// 		Target: to.Ptr(""),
	// 		Details: []*armmanagednetworkfabric.ErrorDetail{
	// 		},
	// 	},
	// 	ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// }
}
