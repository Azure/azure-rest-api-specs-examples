package armmongocluster_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: 2025-09-01/MongoClusters_UserCreateOrUpdate.json
func ExampleUsersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewUsersClient().BeginCreateOrUpdate(ctx, "TestGroup", "myMongoCluster", "uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu", armmongocluster.User{
		Properties: &armmongocluster.UserProperties{
			Roles: []*armmongocluster.DatabaseRole{
				{
					Role: to.Ptr(armmongocluster.UserRoleRoot),
					Db:   to.Ptr("admin"),
				},
			},
			IdentityProvider: &armmongocluster.EntraIdentityProvider{
				Type: to.Ptr(armmongocluster.IdentityProviderTypeMicrosoftEntraID),
				Properties: &armmongocluster.EntraIdentityProviderProperties{
					PrincipalType: to.Ptr(armmongocluster.EntraPrincipalTypeUser),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongocluster.UsersClientCreateOrUpdateResponse{
	// 	User: &armmongocluster.User{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster/users/uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu"),
	// 		Name: to.Ptr("uuuuuuuu-uuuu-uuuu-uuuu-uuuuuuuuuuuu"),
	// 		Type: to.Ptr("/Microsoft.DocumentDB/mongoClusters/users"),
	// 		SystemData: &armmongocluster.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armmongocluster.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		},
	// 		Properties: &armmongocluster.UserProperties{
	// 			ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateInProgress),
	// 			Roles: []*armmongocluster.DatabaseRole{
	// 				{
	// 					Role: to.Ptr(armmongocluster.UserRoleRoot),
	// 					Db: to.Ptr("admin"),
	// 				},
	// 			},
	// 			IdentityProvider: &armmongocluster.EntraIdentityProvider{
	// 				Type: to.Ptr(armmongocluster.IdentityProviderTypeMicrosoftEntraID),
	// 				Properties: &armmongocluster.EntraIdentityProviderProperties{
	// 					PrincipalType: to.Ptr(armmongocluster.EntraPrincipalTypeUser),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
