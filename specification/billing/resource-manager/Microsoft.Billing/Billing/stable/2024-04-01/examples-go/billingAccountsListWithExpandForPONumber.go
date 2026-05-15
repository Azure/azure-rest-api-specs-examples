package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingAccountsListWithExpandForPONumber.json
func ExampleAccountsClient_NewListPager_billingAccountsListWithExpandForPoNumber() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(&armbilling.AccountsClientListOptions{
		Expand: to.Ptr("soldTo,enrollmentDetails/poNumber")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armbilling.AccountsClientListResponse{
		// 	AccountListResult: armbilling.AccountListResult{
		// 		Value: []*armbilling.Account{
		// 			{
		// 				Name: to.Ptr("6564892"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6564892"),
		// 				Properties: &armbilling.AccountProperties{
		// 					AccountStatus: to.Ptr(armbilling.AccountStatusActive),
		// 					AccountSubType: to.Ptr(armbilling.AccountSubTypeNone),
		// 					AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
		// 					AgreementType: to.Ptr(armbilling.AgreementTypeEnterpriseAgreement),
		// 					DisplayName: to.Ptr("Enterprise Account"),
		// 					EnrollmentDetails: &armbilling.AccountPropertiesEnrollmentDetails{
		// 						BillingCycle: to.Ptr("Monthly"),
		// 						Channel: to.Ptr("EaDirect"),
		// 						Cloud: to.Ptr("Azure Commercial"),
		// 						CountryCode: to.Ptr("US"),
		// 						Currency: to.Ptr("USD"),
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28Z"); return t}()),
		// 						ExtendedTermOption: to.Ptr(armbilling.ExtendedTermOptionOptedOut),
		// 						PoNumber: to.Ptr("poNumber123"),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28Z"); return t}()),
		// 						SupportCoverage: to.Ptr("1/26/2021 - 6/30/2021"),
		// 						SupportLevel: to.Ptr(armbilling.SupportLevelStandard),
		// 						Language: to.Ptr("en"),
		// 					},
		// 					HasReadAccess: to.Ptr(true),
		// 					SoldTo: &armbilling.AccountPropertiesSoldTo{
		// 						AddressLine1: to.Ptr("Test Address"),
		// 						City: to.Ptr("City"),
		// 						CompanyName: to.Ptr("Enterprise Company"),
		// 						Country: to.Ptr("US"),
		// 						PostalCode: to.Ptr("00000-1111"),
		// 						Region: to.Ptr("WA"),
		// 					},
		// 				},
		// 				SystemData: &armbilling.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.2606750Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.2606750Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
