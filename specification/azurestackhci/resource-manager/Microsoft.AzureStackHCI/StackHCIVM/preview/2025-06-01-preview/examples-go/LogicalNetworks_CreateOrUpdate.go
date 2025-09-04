package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/LogicalNetworks_CreateOrUpdate.json
func ExampleLogicalNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLogicalNetworksClient().BeginCreateOrUpdate(ctx, "test-rg", "test-lnet", armazurestackhcivm.LogicalNetwork{
		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		},
		Location: to.Ptr("West US2"),
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
	// res = armazurestackhcivm.LogicalNetworksClientCreateOrUpdateResponse{
	// 	LogicalNetwork: &armazurestackhcivm.LogicalNetwork{
	// 		Name: to.Ptr("test-lnet"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/logicalNetworks"),
	// 		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/logicalNetworks/test-lnet"),
	// 		Location: to.Ptr("West US2"),
	// 		Properties: &armazurestackhcivm.LogicalNetworkProperties{
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
	// 		},
	// 	},
	// }
}
