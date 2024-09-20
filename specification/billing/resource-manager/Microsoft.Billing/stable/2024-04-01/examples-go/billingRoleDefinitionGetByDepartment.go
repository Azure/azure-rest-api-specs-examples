package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByDepartment.json
func ExampleRoleDefinitionClient_GetByDepartment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionClient().GetByDepartment(ctx, "123456", "7368531", "50000000-aaaa-bbbb-cccc-100000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleDefinition = armbilling.RoleDefinition{
	// 	Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000000"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/billingRoleDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/departments/7368531/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000000"),
	// 	Properties: &armbilling.RoleDefinitionProperties{
	// 		Permissions: []*armbilling.Permission{
	// 			{
	// 				Actions: []*string{
	// 					to.Ptr("Microsoft.Billing/billingAccounts/departments/read"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/departments/write"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/departments/action"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/departments/delete"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/read"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/write"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/action"),
	// 					to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/delete")},
	// 			}},
	// 			RoleName: to.Ptr("Department Admin"),
	// 		},
	// 	}
}
