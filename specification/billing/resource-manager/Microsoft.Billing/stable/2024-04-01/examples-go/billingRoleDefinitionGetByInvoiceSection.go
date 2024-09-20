package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionGetByInvoiceSection.json
func ExampleRoleDefinitionClient_GetByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionClient().GetByInvoiceSection(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "yyyy-yyyy-yyy-yyy", "30000000-aaaa-bbbb-cccc-100000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleDefinition = armbilling.RoleDefinition{
	// 	Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000000"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections/billingRoleDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"),
	// 	Properties: &armbilling.RoleDefinitionProperties{
	// 		Description: to.Ptr("The Owner role gives the user all permissions including access management rights to the invoice section."),
	// 		Permissions: []*armbilling.Permission{
	// 			{
	// 				Actions: []*string{
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000005"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000006"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000008"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000011"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000012"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000013"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000016"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
	// 					to.Ptr("30000000-aaaa-bbbb-cccc-200000000018")},
	// 			}},
	// 			RoleName: to.Ptr("Invoice section owner"),
	// 		},
	// 	}
}
