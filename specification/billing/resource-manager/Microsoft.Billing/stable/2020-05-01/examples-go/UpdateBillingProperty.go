package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingProperty.json
func ExamplePropertyClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPropertyClient().Update(ctx, armbilling.Property{
		Properties: &armbilling.PropertyProperties{
			CostCenter: to.Ptr("1010"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Property = armbilling.Property{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingProperty"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingProperty/default"),
	// 	Properties: &armbilling.PropertyProperties{
	// 		AccountAdminNotificationEmailAddress: to.Ptr("test@contoso.com"),
	// 		BillingAccountDisplayName: to.Ptr("My Account"),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000"),
	// 		BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 		BillingProfileSpendingLimit: to.Ptr(armbilling.BillingProfileSpendingLimitOn),
	// 		BillingProfileStatus: to.Ptr(armbilling.BillingProfileStatusWarned),
	// 		BillingProfileStatusReasonCode: to.Ptr(armbilling.BillingProfileStatusReasonCodePastDue),
	// 		BillingTenantID: to.Ptr("90000000-0000-0000-0000-000000000011"),
	// 		CostCenter: to.Ptr("1010"),
	// 		InvoiceSectionDisplayName: to.Ptr("Contoso operations invoice section"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/invoiceSections/22000000-0000-0000-0000-000000000000"),
	// 		IsAccountAdmin: to.Ptr(true),
	// 		ProductID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000_00000000-0000-0000-0000-000000000000/invoiceSections/22000000-0000-0000-0000-000000000000/products/44000000-0000-0000-0000-000000000000"),
	// 		ProductName: to.Ptr("Standard Dev"),
	// 		SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
	// 		SKUID: to.Ptr("0001"),
	// 	},
	// }
}
