package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: 2024-07-01-preview/GetDenyAssignmentByNameId.json
func ExampleDenyAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDenyAssignmentsClient().Get(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname", "denyAssignmentId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armauthorization.DenyAssignmentsClientGetResponse{
	// 	DenyAssignment: &armauthorization.DenyAssignment{
	// 		Name: to.Ptr("denyAssignmentId"),
	// 		Type: to.Ptr("Microsoft.Authorization/denyAssignments"),
	// 		ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname/providers/Microsoft.Authorization/denyAssignments/denyAssignmentId"),
	// 		Properties: &armauthorization.DenyAssignmentProperties{
	// 			Description: to.Ptr("Deny assignment description"),
	// 			DenyAssignmentEffect: to.Ptr(armauthorization.DenyAssignmentEffectEnforced),
	// 			DenyAssignmentName: to.Ptr("Deny assignment name"),
	// 			DoNotApplyToChildScopes: to.Ptr(false),
	// 			ExcludePrincipals: []*armauthorization.DenyAssignmentPrincipal{
	// 				{
	// 					Type: to.Ptr("principalType2"),
	// 					ID: to.Ptr("principalId2"),
	// 				},
	// 			},
	// 			IsSystemProtected: to.Ptr(true),
	// 			Permissions: []*armauthorization.DenyAssignmentPermission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("action"),
	// 					},
	// 					DataActions: []*string{
	// 					},
	// 					NotActions: []*string{
	// 					},
	// 					NotDataActions: []*string{
	// 					},
	// 				},
	// 			},
	// 			Principals: []*armauthorization.DenyAssignmentPrincipal{
	// 				{
	// 					Type: to.Ptr("principalType1"),
	// 					ID: to.Ptr("principalId1"),
	// 				},
	// 			},
	// 			Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname"),
	// 		},
	// 		SystemData: &armauthorization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armauthorization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armauthorization.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
