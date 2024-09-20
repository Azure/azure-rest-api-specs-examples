package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentGetByDepartment.json
func ExampleRoleAssignmentsClient_GetByDepartment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().GetByDepartment(ctx, "7898901", "225314", "9dfd08c2-62a3-4d47-85bd-1cdba1408402", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armbilling.RoleAssignment{
	// 	Name: to.Ptr("9dfd08c2-62a3-4d47-85bd-1cdba1408402"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/billingRoleAssignments"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/225314/billingRoleAssignments/9dfd08c2-62a3-4d47-85bd-1cdba1408402"),
	// 	Properties: &armbilling.RoleAssignmentProperties{
	// 		CreatedByUserEmailAddress: to.Ptr("test@billtest900006.onmicrosoft.com"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/225314/billingRoleDefinitions/c15c22c0-9faf-424c-9b7e-bd91c06a240b"),
	// 		Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/225314"),
	// 		UserAuthenticationType: to.Ptr("Organization"),
	// 		UserEmailAddress: to.Ptr("a_owner@billtest900006.onmicrosoft.com"),
	// 	},
	// }
}
