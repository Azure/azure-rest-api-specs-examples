package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionValidateMoveFailure.json
func ExampleSubscriptionsClient_ValidateMoveEligibility_billingSubscriptionValidateMoveFailure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().ValidateMoveEligibility(ctx, "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31", "6b96d3f2-9008-4a9d-912f-f87744185aa3", armbilling.MoveBillingSubscriptionRequest{
		DestinationInvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveBillingSubscriptionEligibilityResult = armbilling.MoveBillingSubscriptionEligibilityResult{
	// 	ErrorDetails: &armbilling.MoveBillingSubscriptionErrorDetails{
	// 		Code: to.Ptr(armbilling.SubscriptionTransferValidationErrorCodeSubscriptionNotActive),
	// 		Message: to.Ptr("Invoice Sections can only be changed for active subscriptions."),
	// 	},
	// }
}
