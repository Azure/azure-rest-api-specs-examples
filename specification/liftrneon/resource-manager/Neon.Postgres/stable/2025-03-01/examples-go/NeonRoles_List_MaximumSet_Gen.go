package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/NeonRoles_List_MaximumSet_Gen.json
func ExampleNeonRolesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNeonRolesClient().NewListPager("rgneon", "test-org", "entity-name", "entity-name", nil)
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
		// page = armneonpostgres.NeonRolesClientListResponse{
		// 	NeonRoleListResult: armneonpostgres.NeonRoleListResult{
		// 		Value: []*armneonpostgres.NeonRole{
		// 			{
		// 				Properties: &armneonpostgres.NeonRoleProperties{
		// 					EntityID: to.Ptr("entity-id"),
		// 					EntityName: to.Ptr("entity-name"),
		// 					CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
		// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
		// 					Attributes: []*armneonpostgres.Attributes{
		// 						{
		// 							Name: to.Ptr("trhvzyvaqy"),
		// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
		// 						},
		// 					},
		// 					BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
		// 					Permissions: []*string{
		// 						to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
		// 					},
		// 					IsSuperUser: to.Ptr(true),
		// 				},
		// 				ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/entity-name/branches/entity-name/neonRoles/entity-name"),
		// 				Name: to.Ptr("tpbnco"),
		// 				Type: to.Ptr("vqfmoiwt"),
		// 				SystemData: &armneonpostgres.SystemData{
		// 					CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
		// 					CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
		// 					LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
