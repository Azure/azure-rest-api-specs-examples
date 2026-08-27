package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-05-01-preview/Administrators_List.json
func ExampleAdministratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdministratorsClient().NewListPager("exampleresourcegroup", "examplecluster", nil)
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
		// page = armhorizondb.AdministratorsClientListResponse{
		// 	AdministratorListResult: armhorizondb.AdministratorListResult{
		// 		Value: []*armhorizondb.Administrator{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/administrators/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 				Name: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/administrators"),
		// 				Properties: &armhorizondb.AdministratorProperties{
		// 					PrincipalName: to.Ptr("admin@contoso.com"),
		// 					PrincipalType: to.Ptr(armhorizondb.PrincipalTypesUser),
		// 					ObjectID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					TenantID: to.Ptr("11111111-2222-3333-4444-555555555555"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/administrators/99999999-8888-7777-6666-555555555555"),
		// 				Name: to.Ptr("99999999-8888-7777-6666-555555555555"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/administrators"),
		// 				Properties: &armhorizondb.AdministratorProperties{
		// 					PrincipalName: to.Ptr("DBA Team"),
		// 					PrincipalType: to.Ptr(armhorizondb.PrincipalTypesGroup),
		// 					ObjectID: to.Ptr("99999999-8888-7777-6666-555555555555"),
		// 					TenantID: to.Ptr("11111111-2222-3333-4444-555555555555"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
