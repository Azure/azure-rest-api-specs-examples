package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/RoleMappings_List_MinimumSet_Gen.json
func ExampleRoleMappingsClient_NewListPager_roleMappingsListMinimumSetListRoleMappingsWithSingleRole() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("1B4766AF-8D4B-4B44-9CF1-587E003DFF7F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleMappingsClient().NewListPager("rgcommvault", "myCloudAccount", nil)
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
		// page = armcommvaultcontentstore.RoleMappingsClientListResponse{
		// 	RoleMappingListResult: armcommvaultcontentstore.RoleMappingListResult{
		// 		Value: []*armcommvaultcontentstore.RoleMapping{
		// 			{
		// 				Properties: &armcommvaultcontentstore.RoleMappingProperties{
		// 					Roles: []*armcommvaultcontentstore.RoleAssignment{
		// 						{
		// 							RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupAdmin),
		// 							Entities: []*armcommvaultcontentstore.EntityInfo{
		// 								{
		// 									ID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
		// 									DisplayName: to.Ptr("Tenant Admins"),
		// 									EntityType: to.Ptr(armcommvaultcontentstore.EntityTypeGroup),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/1B4766AF-8D4B-4B44-9CF1-587E003DFF7F/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/roleMappings/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Commvault.ContentStore/cloudAccounts/roleMappings"),
		// 				SystemData: &armcommvaultcontentstore.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
