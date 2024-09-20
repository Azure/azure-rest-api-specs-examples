package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountsListByDepartment.json
func ExampleEnrollmentAccountsClient_NewListByDepartmentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnrollmentAccountsClient().NewListByDepartmentPager("6564892", "164821", &armbilling.EnrollmentAccountsClientListByDepartmentOptions{Filter: nil,
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  nil,
	})
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
		// page.EnrollmentAccountListResult = armbilling.EnrollmentAccountListResult{
		// 	Value: []*armbilling.EnrollmentAccount{
		// 		{
		// 			Name: to.Ptr("257698"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/enrollmentAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6564892/departments/164821/enrollmentAccounts/257698"),
		// 			Properties: &armbilling.EnrollmentAccountProperties{
		// 				AccountOwner: to.Ptr("account0@contoso.com"),
		// 				AuthType: to.Ptr("MicrosoftAccount"),
		// 				CostCenter: to.Ptr("C0"),
		// 				DepartmentDisplayName: to.Ptr("TestDept"),
		// 				DepartmentID: to.Ptr("164821"),
		// 				DisplayName: to.Ptr("AccountName0"),
		// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 				IsDevTestEnabled: to.Ptr(true),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 				Status: to.Ptr("Active"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("264698"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/enrollmentAccounts"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6564892/departments/164821/enrollmentAccounts/264698"),
		// 			Properties: &armbilling.EnrollmentAccountProperties{
		// 				AccountOwner: to.Ptr("account1@contoso.com"),
		// 				AuthType: to.Ptr("Organization"),
		// 				CostCenter: to.Ptr("C1"),
		// 				DepartmentDisplayName: to.Ptr("TestDept"),
		// 				DepartmentID: to.Ptr("164821"),
		// 				DisplayName: to.Ptr("AccountName1"),
		// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
		// 				IsDevTestEnabled: to.Ptr(false),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 				Status: to.Ptr("Active"),
		// 			},
		// 	}},
		// }
	}
}
