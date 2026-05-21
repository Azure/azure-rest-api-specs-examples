package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2025-09-01/NGroupsList.json
func ExampleNGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNGroupsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerinstance.NGroupsClientListResponse{
		// 	NGroupsListResult: armcontainerinstance.NGroupsListResult{
		// 		Value: []*armcontainerinstance.NGroup{
		// 			{
		// 				Properties: &armcontainerinstance.NGroupProperties{
		// 					UpdateProfile: &armcontainerinstance.UpdateProfile{
		// 						UpdateMode: to.Ptr(armcontainerinstance.NGroupUpdateModeManual),
		// 					},
		// 					ContainerGroupProfiles: []*armcontainerinstance.ContainerGroupProfileStub{
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerinstance.NGroupProvisioningStateSucceeded),
		// 					ElasticProfile: &armcontainerinstance.ElasticProfile{
		// 						MaintainDesiredCount: to.Ptr(true),
		// 						DesiredCount: to.Ptr[int32](1),
		// 					},
		// 				},
		// 				Name: to.Ptr("demo-ngroup"),
		// 				Type: to.Ptr("Microsoft.ContainerInstance/ngroups"),
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.ContainerInstance/ngroups/demo-ngroup"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
