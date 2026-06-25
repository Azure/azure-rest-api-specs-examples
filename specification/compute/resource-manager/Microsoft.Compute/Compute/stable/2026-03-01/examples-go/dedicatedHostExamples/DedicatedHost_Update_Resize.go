package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/dedicatedHostExamples/DedicatedHost_Update_Resize.json
func ExampleDedicatedHostsClient_BeginUpdate_dedicatedHostUpdateResize() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDedicatedHostsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", armcompute.DedicatedHostUpdate{
		SKU: &armcompute.SKU{
			Name: to.Ptr("DSv3-Type1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DedicatedHostsClientUpdateResponse{
	// 	DedicatedHost: armcompute.DedicatedHost{
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armcompute.DedicatedHostProperties{
	// 			PlatformFaultDomain: to.Ptr[int32](1),
	// 			AutoReplaceOnFailure: to.Ptr(true),
	// 			HostID: to.Ptr("{GUID}"),
	// 			VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("aaaa"),
	// 				},
	// 			},
	// 			LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
	// 			ProvisioningState: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("DSv3-Type1"),
	// 			Tier: to.Ptr("aaa"),
	// 			Capacity: to.Ptr[int64](7),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
	// 		Name: to.Ptr("myDedicatedHost"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	},
	// }
}
