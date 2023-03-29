package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleDefinition.json
func ExampleRoleDefinitionsClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionsClient().GetByBillingAccount(ctx, "{billingAccountName}", "{billingRoleDefinitionName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleDefinition = armbilling.RoleDefinition{
	// 	Name: to.Ptr("{billingRoleDefinitionName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/{billingRoleDefinitionName}"),
	// 	Properties: &armbilling.RoleDefinitionProperties{
	// 		Description: to.Ptr("The Owner role gives the user all permissions including access management rights to the billing account."),
	// 		Permissions: []*armbilling.PermissionsProperties{
	// 			{
	// 				Actions: []*string{
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000000"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000008"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000001"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000002"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
	// 					to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
	// 					to.Ptr("20000000-aaaa-bbbb-cccc-200000000000"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
	// 					to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
	// 					to.Ptr("20000000-aaaa-bbbb-cccc-200000000002"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
	// 					to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
	// 					to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
	// 					to.Ptr("40000000-aaaa-bbbb-cccc-200000000008")},
	// 					NotActions: []*string{
	// 					},
	// 			}},
	// 			RoleName: to.Ptr("Billing acount owner"),
	// 		},
	// 	}
}
