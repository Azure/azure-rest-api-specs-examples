package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7678455846b1000fd31db27596e4ca3d299a872/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Addons_CreateOrUpdate_VR.json
func ExampleAddonsClient_BeginCreateOrUpdate_addonsCreateOrUpdateVr() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAddonsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "vr", armavs.Addon{
		Properties: &armavs.AddonVrProperties{
			AddonType: to.Ptr(armavs.AddonTypeVR),
			VrsCount:  to.Ptr[int32](1),
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
	// res.Addon = armavs.Addon{
	// 	Name: to.Ptr("vr"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/addons"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/addons/vr"),
	// 	Properties: &armavs.AddonVrProperties{
	// 		AddonType: to.Ptr(armavs.AddonTypeVR),
	// 		ProvisioningState: to.Ptr(armavs.AddonProvisioningStateSucceeded),
	// 		VrsCount: to.Ptr[int32](1),
	// 	},
	// }
}
