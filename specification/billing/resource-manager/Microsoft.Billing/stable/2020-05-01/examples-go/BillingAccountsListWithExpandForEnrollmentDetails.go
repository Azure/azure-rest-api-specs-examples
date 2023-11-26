package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsListWithExpandForEnrollmentDetails.json
func ExampleAccountsClient_NewListPager_billingAccountsListWithExpandForEnrollmentDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(&armbilling.AccountsClientListOptions{Expand: to.Ptr("enrollmentDetails,departments,enrollmentAccounts")})
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
		// page.AccountListResult = armbilling.AccountListResult{
		// 	Value: []*armbilling.Account{
		// 		{
		// 			Name: to.Ptr("7645820"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/7645820"),
		// 			Properties: &armbilling.AccountProperties{
		// 				AgreementType: to.Ptr(armbilling.AgreementTypeEnterpriseAgreement),
		// 				Departments: []*armbilling.Department{
		// 					{
		// 						Name: to.Ptr("departmentId1"),
		// 						Type: to.Ptr("Microsoft.Billing/departments"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/departments/departmentId1"),
		// 						Properties: &armbilling.DepartmentProperties{
		// 							CostCenter: to.Ptr("C1"),
		// 							DepartmentName: to.Ptr("departmentName1"),
		// 							Status: to.Ptr("Active"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("departmentId2"),
		// 						Type: to.Ptr("Microsoft.Billing/departments"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/departments/departmentId2"),
		// 						Properties: &armbilling.DepartmentProperties{
		// 							CostCenter: to.Ptr("C4"),
		// 							DepartmentName: to.Ptr("departmentName2"),
		// 							Status: to.Ptr("Active"),
		// 						},
		// 				}},
		// 				EnrollmentAccounts: []*armbilling.EnrollmentAccount{
		// 					{
		// 						Name: to.Ptr("accountId0"),
		// 						Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts/accountId0"),
		// 						Properties: &armbilling.EnrollmentAccountProperties{
		// 							AccountName: to.Ptr("AccountName0"),
		// 							CostCenter: to.Ptr("C0"),
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 							Status: to.Ptr("Active"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("accountId1"),
		// 						Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts/accountId1"),
		// 						Properties: &armbilling.EnrollmentAccountProperties{
		// 							AccountName: to.Ptr("AccountName1"),
		// 							CostCenter: to.Ptr("C4"),
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 							Status: to.Ptr("Active"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("accountId2"),
		// 						Type: to.Ptr("Microsoft.Billing/enrollmentAccounts"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts/accountId2"),
		// 						Properties: &armbilling.EnrollmentAccountProperties{
		// 							AccountName: to.Ptr("AccountName2"),
		// 							CostCenter: to.Ptr("C4"),
		// 							EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 							StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 							Status: to.Ptr("Active"),
		// 						},
		// 				}},
		// 				EnrollmentDetails: &armbilling.Enrollment{
		// 					BillingCycle: to.Ptr("Monthly"),
		// 					Channel: to.Ptr("EaDirect"),
		// 					CountryCode: to.Ptr("US"),
		// 					Currency: to.Ptr("USD"),
		// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 					Policies: &armbilling.EnrollmentPolicies{
		// 						AccountOwnerViewCharges: to.Ptr(true),
		// 						DepartmentAdminViewCharges: to.Ptr(true),
		// 						MarketplaceEnabled: to.Ptr(true),
		// 						ReservedInstancesEnabled: to.Ptr(true),
		// 					},
		// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 					Status: to.Ptr("Active"),
		// 					Language: to.Ptr("en"),
		// 				},
		// 				SoldTo: &armbilling.AddressDetails{
		// 					AddressLine1: to.Ptr("Test Address"),
		// 					AddressLine2: to.Ptr("Test Address"),
		// 					AddressLine3: to.Ptr("Test Address"),
		// 					City: to.Ptr("City"),
		// 					Country: to.Ptr("US"),
		// 					PostalCode: to.Ptr("00000"),
		// 					Region: to.Ptr("WA"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
