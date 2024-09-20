package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByBillingProfile.json
func ExamplePoliciesClient_BeginCreateOrUpdateByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoliciesClient().BeginCreateOrUpdateByBillingProfile(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", armbilling.ProfilePolicy{
		Properties: &armbilling.ProfilePolicyProperties{
			InvoiceSectionLabelManagement: to.Ptr(armbilling.InvoiceSectionLabelManagementPolicyAllowed),
			MarketplacePurchases:          to.Ptr(armbilling.MarketplacePurchasesPolicyAllAllowed),
			ReservationPurchases:          to.Ptr(armbilling.ReservationPurchasesPolicyAllowed),
			SavingsPlanPurchases:          to.Ptr(armbilling.SavingsPlanPurchasesPolicyAllowed),
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
	// res.ProfilePolicy = armbilling.ProfilePolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/policies"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/xxxx-xxxx-xxx-xxx/policies/default"),
	// 	Properties: &armbilling.ProfilePolicyProperties{
	// 		InvoiceSectionLabelManagement: to.Ptr(armbilling.InvoiceSectionLabelManagementPolicyAllowed),
	// 		MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyAllAllowed),
	// 		ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyAllowed),
	// 		SavingsPlanPurchases: to.Ptr(armbilling.SavingsPlanPurchasesPolicyAllowed),
	// 	},
	// }
}
