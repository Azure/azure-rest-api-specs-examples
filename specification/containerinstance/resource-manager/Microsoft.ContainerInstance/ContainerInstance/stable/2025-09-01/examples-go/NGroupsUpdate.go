package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/NGroupsUpdate.json
func ExampleNGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNGroupsClient().BeginUpdate(ctx, "demo", "demo-ngroup", armcontainerinstance.NGroupPatch{
		Tags: map[string]*string{
			"env": to.Ptr("test"),
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
	// res = armcontainerinstance.NGroupsClientUpdateResponse{
	// 	NGroup: armcontainerinstance.NGroup{
	// 		Properties: &armcontainerinstance.NGroupProperties{
	// 			UpdateProfile: &armcontainerinstance.UpdateProfile{
	// 				UpdateMode: to.Ptr(armcontainerinstance.NGroupUpdateModeManual),
	// 			},
	// 			ContainerGroupProfiles: []*armcontainerinstance.ContainerGroupProfileStub{
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerinstance.NGroupProvisioningStateSucceeded),
	// 			ElasticProfile: &armcontainerinstance.ElasticProfile{
	// 				MaintainDesiredCount: to.Ptr(true),
	// 				DesiredCount: to.Ptr[int32](1),
	// 			},
	// 		},
	// 		Name: to.Ptr("demo-ngroup"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/ngroups"),
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.ContainerInstance/ngroups/demo-ngroup"),
	// 	},
	// }
}
