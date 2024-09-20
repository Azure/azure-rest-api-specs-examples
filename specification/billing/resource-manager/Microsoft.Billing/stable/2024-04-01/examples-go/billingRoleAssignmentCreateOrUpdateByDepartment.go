package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateOrUpdateByDepartment.json
func ExampleRoleAssignmentsClient_BeginCreateOrUpdateByDepartment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoleAssignmentsClient().BeginCreateOrUpdateByDepartment(ctx, "7898901", "12345", "9dfd08c2-62a3-4d47-85bd-1cdba1408402", armbilling.RoleAssignment{
		Properties: &armbilling.RoleAssignmentProperties{
			PrincipalID:       to.Ptr("00000000-0000-0000-0000-000000000000"),
			PrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
			RoleDefinitionID:  to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/12345/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db"),
			UserEmailAddress:  to.Ptr("john@contoso.com"),
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
	// res.RoleAssignment = armbilling.RoleAssignment{
	// 	Name: to.Ptr("9dfd08c2-62a3-4d47-85bd-1cdba1408402"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/billingRoleAssignments"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/12345/billingRoleAssignments/9dfd08c2-62a3-4d47-85bd-1cdba1408402"),
	// 	Properties: &armbilling.RoleAssignmentProperties{
	// 		CreatedByPrincipalID: to.Ptr("60d97094-2be4-46cc-a4fe-3633021a25b9"),
	// 		CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		PrincipalTenantID: to.Ptr("7ca289b9-c32d-4f01-8566-7ff93261d76f"),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/12345/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db"),
	// 		Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7898901/departments/12345"),
	// 	},
	// }
}
