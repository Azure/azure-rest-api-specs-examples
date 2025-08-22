package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba99beec358a40ee08dae7f12f6a989aad6ce6d1/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Subvolumes_Create.json
func ExampleSubvolumesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubvolumesClient().BeginCreate(ctx, "myRG", "account1", "pool1", "volume1", "subvolume1", armnetapp.SubvolumeInfo{
		Properties: &armnetapp.SubvolumeProperties{
			Path: to.Ptr("/subvolumePath"),
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
	// res.SubvolumeInfo = armnetapp.SubvolumeInfo{
	// 	Name: to.Ptr("account1/pool1/volume1/subvolume1"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/subvolumes/subvolume1"),
	// 	Properties: &armnetapp.SubvolumeProperties{
	// 		Path: to.Ptr("/subvolumePath"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Size: to.Ptr[int64](0),
	// 	},
	// }
}
