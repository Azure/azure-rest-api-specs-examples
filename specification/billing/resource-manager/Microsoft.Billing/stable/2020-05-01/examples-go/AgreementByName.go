package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AgreementByName.json
func ExampleAgreementsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgreementsClient().Get(ctx, "{billingAccountName}", "{agreementName}", &armbilling.AgreementsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agreement = armbilling.Agreement{
	// 	Name: to.Ptr("{agreementName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/agreements"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements/{agreementName}"),
	// 	Properties: &armbilling.AgreementProperties{
	// 		AcceptanceMode: to.Ptr(armbilling.AcceptanceModeClickToAccept),
	// 		AgreementLink: to.Ptr("https://agreementuri1.com"),
	// 		Category: to.Ptr(armbilling.CategoryMicrosoftCustomerAgreement),
	// 		EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-05T00:00:00.000Z"); return t}()),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-05T00:00:00.000Z"); return t}()),
	// 		Participants: []*armbilling.Participants{
	// 			{
	// 				Email: to.Ptr("abc@contoso.com"),
	// 				Status: to.Ptr("Accepted"),
	// 				StatusDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-01T00:00:00.000Z"); return t}()),
	// 			},
	// 			{
	// 				Email: to.Ptr("xyz@contoso.com"),
	// 				Status: to.Ptr("Declined"),
	// 				StatusDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-02T00:00:00.000Z"); return t}()),
	// 		}},
	// 		Status: to.Ptr("Published"),
	// 	},
	// }
}
