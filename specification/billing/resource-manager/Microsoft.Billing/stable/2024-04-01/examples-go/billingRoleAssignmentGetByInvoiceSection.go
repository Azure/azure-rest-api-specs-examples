package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentGetByInvoiceSection.json
func ExampleRoleAssignmentsClient_GetByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().GetByInvoiceSection(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB", "xxxx-xxxx-xxx-xxx", "10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armbilling.RoleAssignment{
	// 	Name: to.Ptr("10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections/billingRoleAssignments"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/invoiceSections/xxxx-xxxx-xxx-xxx/billingRoleAssignments/10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 	Properties: &armbilling.RoleAssignmentProperties{
	// 		CreatedByPrincipalID: to.Ptr("46b831ec-42b2-4f1a-8b54-3fd5ff9d6aa1"),
	// 		CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/invoiceSections/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000000"),
	// 		Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/invoiceSections/xxxx-xxxx-xxx-xxx"),
	// 	},
	// }
}
