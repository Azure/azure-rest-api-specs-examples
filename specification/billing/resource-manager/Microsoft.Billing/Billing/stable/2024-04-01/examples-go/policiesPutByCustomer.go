package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/policiesPutByCustomer.json
func ExamplePoliciesClient_BeginCreateOrUpdateByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoliciesClient().BeginCreateOrUpdateByCustomer(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "11111111-1111-1111-1111-111111111111", armbilling.CustomerPolicy{
		Properties: &armbilling.CustomerPolicyProperties{
			ViewCharges: to.Ptr(armbilling.ViewChargesPolicyAllowed),
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
	// res = armbilling.PoliciesClientCreateOrUpdateByCustomerResponse{
	// 	CustomerPolicy: armbilling.CustomerPolicy{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/policies"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111/policies/default"),
	// 		Properties: &armbilling.CustomerPolicyProperties{
	// 			ViewCharges: to.Ptr(armbilling.ViewChargesPolicyAllowed),
	// 		},
	// 	},
	// }
}
