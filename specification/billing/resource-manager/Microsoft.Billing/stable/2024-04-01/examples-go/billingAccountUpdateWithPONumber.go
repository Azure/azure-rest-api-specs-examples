package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountUpdateWithPONumber.json
func ExampleAccountsClient_BeginUpdate_billingAccountUpdateWithPoNumber() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "6575495", armbilling.AccountPatch{
		Properties: &armbilling.AccountProperties{
			EnrollmentDetails: &armbilling.AccountPropertiesEnrollmentDetails{
				PoNumber: to.Ptr("poNumber123"),
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
	// 	Name: to.Ptr("6575495"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6575495"),
	// 	SystemData: &armbilling.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.260Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 	},
	// 	Properties: &armbilling.AccountProperties{
	// 		AccountStatus: to.Ptr(armbilling.AccountStatusActive),
	// 		AccountSubType: to.Ptr(armbilling.AccountSubTypeNone),
	// 		AccountType: to.Ptr(armbilling.AccountTypeEnterprise),
	// 		AgreementType: to.Ptr(armbilling.AgreementTypeEnterpriseAgreement),
	// 		DisplayName: to.Ptr("Enterprise Account"),
	// 		EnrollmentDetails: &armbilling.AccountPropertiesEnrollmentDetails{
	// 			BillingCycle: to.Ptr("Monthly"),
	// 			Channel: to.Ptr("EaDirect"),
	// 			Cloud: to.Ptr("Azure Commercial"),
	// 			CountryCode: to.Ptr("US"),
	// 			Currency: to.Ptr("USD"),
	// 			EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
	// 			ExtendedTermOption: to.Ptr(armbilling.ExtendedTermOptionOptedOut),
	// 			PoNumber: to.Ptr("poNumber123"),
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
	// 			SupportCoverage: to.Ptr("1/26/2021 - 6/30/2021"),
	// 			SupportLevel: to.Ptr(armbilling.SupportLevelStandard),
	// 			Language: to.Ptr("en"),
	// 		},
	// 		HasReadAccess: to.Ptr(true),
	// 		SoldTo: &armbilling.AccountPropertiesSoldTo{
	// 			AddressLine1: to.Ptr("Test Address"),
	// 			AddressLine2: to.Ptr("Test Address"),
	// 			AddressLine3: to.Ptr("Test Address"),
	// 			City: to.Ptr("City"),
	// 			Country: to.Ptr("US"),
	// 			PostalCode: to.Ptr("00000"),
	// 			Region: to.Ptr("WA"),
	// 		},
	// 	},
	// }
}
