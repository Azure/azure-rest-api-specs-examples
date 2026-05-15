package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/policiesGetByBillingAccount.json
func ExamplePoliciesClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoliciesClient().GetByBillingAccount(ctx, "1234567", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.PoliciesClientGetByBillingAccountResponse{
	// 	AccountPolicy: armbilling.AccountPolicy{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/policies"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234567/policies/default"),
	// 		Properties: &armbilling.AccountPolicyProperties{
	// 			EnterpriseAgreementPolicies: &armbilling.AccountPolicyPropertiesEnterpriseAgreementPolicies{
	// 				AuthenticationType: to.Ptr(armbilling.EnrollmentAuthLevelStateOrganizationalAccountOnly),
	// 			},
	// 			MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyAllAllowed),
	// 			ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyAllowed),
	// 			SavingsPlanPurchases: to.Ptr(armbilling.SavingsPlanPurchasesPolicyNotAllowed),
	// 		},
	// 	},
	// }
}
