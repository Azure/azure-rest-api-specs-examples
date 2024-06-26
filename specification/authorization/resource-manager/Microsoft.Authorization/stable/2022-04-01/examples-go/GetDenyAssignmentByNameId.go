package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetDenyAssignmentByNameId.json
func ExampleDenyAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDenyAssignmentsClient().Get(ctx, "subscriptions/subId/resourcegroups/rgname", "denyAssignmentId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DenyAssignment = armauthorization.DenyAssignment{
	// 	Name: to.Ptr("denyAssignmentId"),
	// 	Type: to.Ptr("Microsoft.Authorization/denyAssignments"),
	// 	ID: to.Ptr("/subscriptions/subId/resourcegroups/rgname/providers/Microsoft.Authorization/denyAssignments/denyAssignmentId"),
	// 	Properties: &armauthorization.DenyAssignmentProperties{
	// 		Description: to.Ptr("Deny assignment description"),
	// 		DenyAssignmentName: to.Ptr("Deny assignment name"),
	// 		DoNotApplyToChildScopes: to.Ptr(false),
	// 		ExcludePrincipals: []*armauthorization.Principal{
	// 			{
	// 				Type: to.Ptr("principalType2"),
	// 				ID: to.Ptr("principalId2"),
	// 		}},
	// 		IsSystemProtected: to.Ptr(true),
	// 		Permissions: []*armauthorization.DenyAssignmentPermission{
	// 			{
	// 				Actions: []*string{
	// 					to.Ptr("action")},
	// 					DataActions: []*string{
	// 					},
	// 					NotActions: []*string{
	// 					},
	// 					NotDataActions: []*string{
	// 					},
	// 			}},
	// 			Principals: []*armauthorization.Principal{
	// 				{
	// 					Type: to.Ptr("principalType1"),
	// 					ID: to.Ptr("principalId1"),
	// 			}},
	// 			Scope: to.Ptr("/subscriptions/subId/resourcegroups/rgname"),
	// 		},
	// 	}
}
