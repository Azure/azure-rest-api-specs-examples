package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: 2024-07-01-preview/DenyAssignments_CreateForSubscription.json
func ExampleDenyAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDenyAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", "64b75d79-7a26-4341-944e-4f1a19f0e6ca", armauthorization.DenyAssignment{
		Properties: &armauthorization.DenyAssignmentProperties{
			Description:             to.Ptr("Prevent all users from deleting critical resources in the subscription."),
			DenyAssignmentEffect:    to.Ptr(armauthorization.DenyAssignmentEffectEnforced),
			DenyAssignmentName:      to.Ptr("Deny delete on critical resources"),
			DoNotApplyToChildScopes: to.Ptr(false),
			ExcludePrincipals: []*armauthorization.DenyAssignmentPrincipal{
				{
					Type: to.Ptr("ServicePrincipal"),
					ID:   to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
				},
			},
			Permissions: []*armauthorization.DenyAssignmentPermission{
				{
					Actions: []*string{
						to.Ptr("*/delete"),
					},
					DataActions:    []*string{},
					NotActions:     []*string{},
					NotDataActions: []*string{},
				},
			},
			Principals: []*armauthorization.DenyAssignmentPrincipal{
				{
					Type: to.Ptr("SystemDefined"),
					ID:   to.Ptr("00000000-0000-0000-0000-000000000000"),
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
	// res = armauthorization.DenyAssignmentsClientCreateOrUpdateResponse{
	// 	DenyAssignment: &armauthorization.DenyAssignment{
	// 		Name: to.Ptr("64b75d79-7a26-4341-944e-4f1a19f0e6ca"),
	// 		Type: to.Ptr("Microsoft.Authorization/denyAssignments"),
	// 		ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/denyAssignments/64b75d79-7a26-4341-944e-4f1a19f0e6ca"),
	// 		Properties: &armauthorization.DenyAssignmentProperties{
	// 			Description: to.Ptr("Prevent all users from deleting critical resources in the subscription."),
	// 			CreatedBy: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
	// 			DenyAssignmentEffect: to.Ptr(armauthorization.DenyAssignmentEffectEnforced),
	// 			DenyAssignmentName: to.Ptr("Deny delete on critical resources"),
	// 			DoNotApplyToChildScopes: to.Ptr(false),
	// 			ExcludePrincipals: []*armauthorization.DenyAssignmentPrincipal{
	// 				{
	// 					Type: to.Ptr("ServicePrincipal"),
	// 					ID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 				},
	// 			},
	// 			IsSystemProtected: to.Ptr(false),
	// 			Permissions: []*armauthorization.DenyAssignmentPermission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("*/delete"),
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
	// 					Type: to.Ptr("SystemDefined"),
	// 					ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 			Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 			UpdatedBy: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 			UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00.0000000Z"); return t}()),
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
