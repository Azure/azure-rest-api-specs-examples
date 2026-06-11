package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkManagerPatch.json
func ExampleManagersClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().Patch(ctx, "rg1", "testNetworkManager", armnetwork.PatchObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ManagersClientPatchResponse{
	// 	Manager: armnetwork.Manager{
	// 		Name: to.Ptr("testNetworkManager"),
	// 		Type: to.Ptr("Microsoft.Network/networkManager"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.ManagerProperties{
	// 			Description: to.Ptr("My Test Network Manager"),
	// 			NetworkManagerScopeAccesses: []*armnetwork.ConfigurationType{
	// 				to.Ptr(armnetwork.ConfigurationTypeConnectivity),
	// 			},
	// 			NetworkManagerScopes: &armnetwork.ManagerPropertiesNetworkManagerScopes{
	// 				ManagementGroups: []*string{
	// 				},
	// 				Subscriptions: []*string{
	// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
