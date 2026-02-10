package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerList.json
func ExampleManagersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagersClient().NewListPager("rg1", &armnetwork.ManagersClientListOptions{Top: nil,
		SkipToken: nil,
	})
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
		// page.ManagerListResult = armnetwork.ManagerListResult{
		// 	Value: []*armnetwork.Manager{
		// 		{
		// 			Name: to.Ptr("testNetworkManager"),
		// 			Type: to.Ptr("Microsoft.Network/networkManagers"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager"),
		// 			Etag: to.Ptr("sadf-asdf-asdf-asdf"),
		// 			Properties: &armnetwork.ManagerProperties{
		// 				Description: to.Ptr("My Test Network Manager"),
		// 				NetworkManagerScopeAccesses: []*armnetwork.ConfigurationType{
		// 					to.Ptr(armnetwork.ConfigurationTypeConnectivity)},
		// 					NetworkManagerScopes: &armnetwork.ManagerPropertiesNetworkManagerScopes{
		// 						ManagementGroups: []*string{
		// 						},
		// 						Subscriptions: []*string{
		// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000")},
		// 						},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 					SystemData: &armnetwork.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 						CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 						CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 						LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					},
		// 			}},
		// 		}
	}
}
