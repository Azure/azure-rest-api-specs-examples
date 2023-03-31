package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsGet.json
func ExampleRoleAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Get(ctx, "TestHubRG", "sdkTestHub", "assignmentName8976", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignmentResourceFormat = armcustomerinsights.RoleAssignmentResourceFormat{
	// 	Name: to.Ptr("azSdkTestHub/assignmentName8976"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/RoleAssignments"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/RoleAssignments/assignmentName8976"),
	// 	Properties: &armcustomerinsights.RoleAssignment{
	// 		AssignmentName: to.Ptr("assignmentName8976"),
	// 		Principals: []*armcustomerinsights.AssignmentPrincipal{
	// 			{
	// 				PrincipalID: to.Ptr("4c54c38f-fa9b-416b-a5a6-d6c8a20cbe7e"),
	// 				PrincipalType: to.Ptr("User"),
	// 			},
	// 			{
	// 				PrincipalID: to.Ptr("93061d15-a505-4f2b-9948-ae25724cf9d5"),
	// 				PrincipalType: to.Ptr("User"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 		Role: to.Ptr(armcustomerinsights.RoleTypesAdmin),
	// 		TenantID: to.Ptr("azsdktesthub"),
	// 	},
	// }
}
