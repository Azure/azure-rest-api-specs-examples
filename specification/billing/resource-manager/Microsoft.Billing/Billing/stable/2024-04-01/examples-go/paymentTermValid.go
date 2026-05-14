package armbilling_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/paymentTermValid.json
func ExampleAccountsClient_ValidatePaymentTerms_paymentTermValid() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ValidatePaymentTerms(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", []*armbilling.PaymentTerm{
		{
			EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-25T22:39:34.2606750Z"); return t }()),
			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.2606750Z"); return t }()),
			Term:      to.Ptr("net10"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.AccountsClientValidatePaymentTermsResponse{
	// 	PaymentTermsEligibilityResult: armbilling.PaymentTermsEligibilityResult{
	// 		EligibilityStatus: to.Ptr(armbilling.PaymentTermsEligibilityStatusValid),
	// 	},
	// }
}
