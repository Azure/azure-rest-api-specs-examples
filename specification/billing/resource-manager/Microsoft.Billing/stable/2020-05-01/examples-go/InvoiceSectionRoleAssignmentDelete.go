package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignmentDelete.json
func ExampleRoleAssignmentsClient_DeleteByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().DeleteByInvoiceSection(ctx, "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", "{billingRoleAssignmentName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armbilling.RoleAssignment{
	// 	Name: to.Ptr("10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections/billingRoleAssignments"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 	Properties: &armbilling.RoleAssignmentProperties{
	// 		CreatedByPrincipalID: to.Ptr("10000000-aaaa-bbbb-cccc-3fd5ff9d6aa1"),
	// 		CreatedByPrincipalTenantID: to.Ptr("7ca289b9-c32d-4f01-8566-7ff93261d76f"),
	// 		CreatedOn: to.Ptr("2018-06-21T21:34:12.2363515+00:00"),
	// 		PrincipalID: to.Ptr("6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 		PrincipalTenantID: to.Ptr("10000000-aaaa-bbbb-cccc-2d7cd011db47"),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000002"),
	// 		Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"),
	// 	},
	// }
}
