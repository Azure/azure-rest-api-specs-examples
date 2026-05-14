package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/policiesPutByBillingAccount.json
func ExamplePoliciesClient_BeginCreateOrUpdateByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoliciesClient().BeginCreateOrUpdateByBillingAccount(ctx, "1234567", armbilling.AccountPolicy{
		Properties: &armbilling.AccountPolicyProperties{
			EnterpriseAgreementPolicies: &armbilling.AccountPolicyPropertiesEnterpriseAgreementPolicies{
				AuthenticationType: to.Ptr(armbilling.EnrollmentAuthLevelStateOrganizationalAccountOnly),
			},
			MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyAllAllowed),
			ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyAllowed),
			SavingsPlanPurchases: to.Ptr(armbilling.SavingsPlanPurchasesPolicyNotAllowed),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.PoliciesClientCreateOrUpdateByBillingAccountResponse{
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
