package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/MicrosoftEntraAdministrators_List.json
func ExampleMicrosoftEntraAdministratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMicrosoftEntraAdministratorsClient().NewListPager("example-resource-group", "example-cluster", nil)
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
		// page = armhorizondb.MicrosoftEntraAdministratorsClientListResponse{
		// 	MicrosoftEntraAdministratorListResult: armhorizondb.MicrosoftEntraAdministratorListResult{
		// 		Value: []*armhorizondb.MicrosoftEntraAdministrator{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/microsoftEntraAdministrators/33333333-3333-3333-3333-333333333333"),
		// 				Name: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/microsoftEntraAdministrators"),
		// 				Properties: &armhorizondb.MicrosoftEntraAdministratorProperties{
		// 					PrincipalName: to.Ptr("admin@contoso.com"),
		// 					PrincipalType: to.Ptr(armhorizondb.PrincipalTypesUser),
		// 					ObjectID: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 					TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/microsoftEntraAdministrators/44444444-4444-4444-4444-444444444444"),
		// 				Name: to.Ptr("44444444-4444-4444-4444-444444444444"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/microsoftEntraAdministrators"),
		// 				Properties: &armhorizondb.MicrosoftEntraAdministratorProperties{
		// 					PrincipalName: to.Ptr("horizondb-automation@contoso.com"),
		// 					PrincipalType: to.Ptr(armhorizondb.PrincipalTypesServicePrincipal),
		// 					ObjectID: to.Ptr("44444444-4444-4444-4444-444444444444"),
		// 					TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
