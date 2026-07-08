package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/RoleMappings_CreateOrUpdate_MaximumSet_Gen.json
func ExampleRoleMappingsClient_CreateOrUpdate_roleMappingsCreateOrUpdateMaximumSetGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("1B4766AF-8D4B-4B44-9CF1-587E003DFF7F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleMappingsClient().CreateOrUpdate(ctx, "rgcommvault", "myCloudAccount", armcommvaultcontentstore.RoleMapping{
		Properties: &armcommvaultcontentstore.RoleMappingProperties{
			Roles: []*armcommvaultcontentstore.RoleAssignment{
				{
					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupAdmin),
					Entities: []*armcommvaultcontentstore.EntityInfo{
						{
							ID:          to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
							DisplayName: to.Ptr("Backup Admins Group"),
							EntityType:  to.Ptr(armcommvaultcontentstore.EntityTypeGroup),
						},
					},
				},
				{
					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupOperator),
					Entities: []*armcommvaultcontentstore.EntityInfo{
						{
							ID:          to.Ptr("11111111-2222-3333-4444-555555555555"),
							DisplayName: to.Ptr("Commvault-Operators"),
							EntityType:  to.Ptr(armcommvaultcontentstore.EntityTypeGroup),
						},
						{
							ID:          to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
							DisplayName: to.Ptr("John Smith"),
							EntityType:  to.Ptr(armcommvaultcontentstore.EntityTypeUser),
						},
					},
				},
				{
					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupUser),
					Entities: []*armcommvaultcontentstore.EntityInfo{
						{
							ID:          to.Ptr("22222222-3333-4444-5555-666666666666"),
							DisplayName: to.Ptr("Jane Doe"),
							EntityType:  to.Ptr(armcommvaultcontentstore.EntityTypeUser),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.RoleMappingsClientCreateOrUpdateResponse{
	// 	RoleMapping: armcommvaultcontentstore.RoleMapping{
	// 		Properties: &armcommvaultcontentstore.RoleMappingProperties{
	// 			Roles: []*armcommvaultcontentstore.RoleAssignment{
	// 				{
	// 					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupAdmin),
	// 					Entities: []*armcommvaultcontentstore.EntityInfo{
	// 						{
	// 							ID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 							DisplayName: to.Ptr("Backup Admins Group"),
	// 							EntityType: to.Ptr(armcommvaultcontentstore.EntityTypeGroup),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupOperator),
	// 					Entities: []*armcommvaultcontentstore.EntityInfo{
	// 						{
	// 							ID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 							DisplayName: to.Ptr("Commvault-Operators"),
	// 							EntityType: to.Ptr(armcommvaultcontentstore.EntityTypeGroup),
	// 						},
	// 						{
	// 							ID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 							DisplayName: to.Ptr("John Smith"),
	// 							EntityType: to.Ptr(armcommvaultcontentstore.EntityTypeUser),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					RoleName: to.Ptr(armcommvaultcontentstore.RoleNameBackupUser),
	// 					Entities: []*armcommvaultcontentstore.EntityInfo{
	// 						{
	// 							ID: to.Ptr("22222222-3333-4444-5555-666666666666"),
	// 							DisplayName: to.Ptr("Jane Doe"),
	// 							EntityType: to.Ptr(armcommvaultcontentstore.EntityTypeUser),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/1B4766AF-8D4B-4B44-9CF1-587E003DFF7F/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount/roleMappings/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Commvault.ContentStore/cloudAccounts/roleMappings"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
	// 		},
	// 	},
	// }
}
