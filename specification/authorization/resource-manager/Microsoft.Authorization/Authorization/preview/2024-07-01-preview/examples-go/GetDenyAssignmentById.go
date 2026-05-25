package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: 2024-07-01-preview/GetDenyAssignmentById.json
func ExampleDenyAssignmentsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDenyAssignmentsClient().GetByID(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname/providers/Microsoft.Authorization/denyAssignments/denyAssignmentId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armauthorization.DenyAssignmentsClientGetByIDResponse{
	// 	DenyAssignment: &armauthorization.DenyAssignment{
	// 		Properties: &armauthorization.DenyAssignmentProperties{
	// 			DenyAssignmentName: to.Ptr("Deny assignment name"),
	// 			Description: to.Ptr("Deny assignment description"),
	// 			Permissions: []*armauthorization.DenyAssignmentPermission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("action"),
	// 					},
	// 					NotActions: []*string{
	// 					},
	// 					DataActions: []*string{
	// 					},
	// 					NotDataActions: []*string{
	// 					},
	// 				},
	// 			},
	// 			Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname"),
	// 			DoNotApplyToChildScopes: to.Ptr(false),
	// 			Principals: []*armauthorization.DenyAssignmentPrincipal{
	// 				{
	// 					ID: to.Ptr("principalId1"),
	// 					Type: to.Ptr("principalType1"),
	// 				},
	// 			},
	// 			ExcludePrincipals: []*armauthorization.DenyAssignmentPrincipal{
	// 				{
	// 					ID: to.Ptr("principalId2"),
	// 					Type: to.Ptr("principalType2"),
	// 				},
	// 			},
	// 			IsSystemProtected: to.Ptr(true),
	// 			DenyAssignmentEffect: to.Ptr(armauthorization.DenyAssignmentEffectEnforced),
	// 		},
	// 		ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/rgname/providers/Microsoft.Authorization/denyAssignments/denyAssignmentId"),
	// 		Type: to.Ptr("Microsoft.Authorization/denyAssignments"),
	// 		Name: to.Ptr("denyAssignmentId"),
	// 		SystemData: &armauthorization.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armauthorization.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armauthorization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
	// 		},
	// 	},
	// }
}
