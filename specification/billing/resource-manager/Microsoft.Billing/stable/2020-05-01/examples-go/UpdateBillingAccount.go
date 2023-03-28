package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingAccount.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "{billingAccountName}", armbilling.AccountUpdateRequest{
		Properties: &armbilling.AccountProperties{
			DisplayName: to.Ptr("Test Account"),
			SoldTo: &armbilling.AddressDetails{
				AddressLine1: to.Ptr("Test Address 1"),
				City:         to.Ptr("Redmond"),
				CompanyName:  to.Ptr("Contoso"),
				Country:      to.Ptr("US"),
				FirstName:    to.Ptr("Test"),
				LastName:     to.Ptr("User"),
				PostalCode:   to.Ptr("12345"),
				Region:       to.Ptr("WA"),
			},
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
	// res.Account = armbilling.Account{
	// 	Name: to.Ptr("{billingAccountName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}"),
	// 	Properties: &armbilling.AccountProperties{
	// 		AccountStatus: to.Ptr(armbilling.AccountStatusActive),
	// 		AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
	// 		AgreementType: to.Ptr(armbilling.AgreementTypeMicrosoftCustomerAgreement),
	// 		DisplayName: to.Ptr("Test Account"),
	// 		HasReadAccess: to.Ptr(true),
	// 		SoldTo: &armbilling.AddressDetails{
	// 			AddressLine1: to.Ptr("Test Address 1"),
	// 			City: to.Ptr("Redmond"),
	// 			CompanyName: to.Ptr("Contoso"),
	// 			Country: to.Ptr("US"),
	// 			FirstName: to.Ptr("Test"),
	// 			LastName: to.Ptr("User"),
	// 			PostalCode: to.Ptr("12345"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 	},
	// }
}
