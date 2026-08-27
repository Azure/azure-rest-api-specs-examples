package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-05-01-preview/Administrators_Get.json
func ExampleAdministratorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdministratorsClient().Get(ctx, "exampleresourcegroup", "examplecluster", "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.AdministratorsClientGetResponse{
	// 	Administrator: armhorizondb.Administrator{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/administrators/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		Name: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/administrators"),
	// 		Properties: &armhorizondb.AdministratorProperties{
	// 			PrincipalName: to.Ptr("admin@contoso.com"),
	// 			PrincipalType: to.Ptr(armhorizondb.PrincipalTypesUser),
	// 			ObjectID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			TenantID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 		},
	// 	},
	// }
}
