package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByBillingAccount.json
func ExamplePoliciesClient_BeginCreateOrUpdateByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountPolicy = armbilling.AccountPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/policies"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234567/policies/default"),
	// 	Properties: &armbilling.AccountPolicyProperties{
	// 		EnterpriseAgreementPolicies: &armbilling.AccountPolicyPropertiesEnterpriseAgreementPolicies{
	// 			AuthenticationType: to.Ptr(armbilling.EnrollmentAuthLevelStateOrganizationalAccountOnly),
	// 		},
	// 		MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyAllAllowed),
	// 		ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyAllowed),
	// 		SavingsPlanPurchases: to.Ptr(armbilling.SavingsPlanPurchasesPolicyNotAllowed),
	// 	},
	// }
}
