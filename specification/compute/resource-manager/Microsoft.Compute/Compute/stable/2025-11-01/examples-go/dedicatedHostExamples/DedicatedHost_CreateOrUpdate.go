package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/dedicatedHostExamples/DedicatedHost_CreateOrUpdate.json
func ExampleDedicatedHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDedicatedHostGroup", "myDedicatedHost", armcompute.DedicatedHost{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"department": to.Ptr("HR"),
		},
		Properties: &armcompute.DedicatedHostProperties{
			PlatformFaultDomain: to.Ptr[int32](1),
		},
		SKU: &armcompute.SKU{
			Name: to.Ptr("DSv3-Type1"),
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
	// res = armcompute.DedicatedHostsClientCreateOrUpdateResponse{
	// 	DedicatedHost: &armcompute.DedicatedHost{
	// 		Properties: &armcompute.DedicatedHostProperties{
	// 			PlatformFaultDomain: to.Ptr[int32](1),
	// 			AutoReplaceOnFailure: to.Ptr(false),
	// 			LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 			HostID: to.Ptr("{GUID}"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		Name: to.Ptr("myDedicatedHost"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("DSv3-Type1"),
	// 		},
	// 	},
	// }
}
