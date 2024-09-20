package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountByDepartment.json
func ExampleEnrollmentAccountsClient_GetByDepartment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnrollmentAccountsClient().GetByDepartment(ctx, "6564892", "164821", "257698", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnrollmentAccount = armbilling.EnrollmentAccount{
	// 	Name: to.Ptr("257698"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/departments/enrollmentAccounts"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6564892/departments/164821/enrollmentAccounts/257698"),
	// 	Properties: &armbilling.EnrollmentAccountProperties{
	// 		AccountOwner: to.Ptr("account0@contoso.com"),
	// 		AuthType: to.Ptr("MicrosoftAccount"),
	// 		CostCenter: to.Ptr("C0"),
	// 		DepartmentDisplayName: to.Ptr("TestDept"),
	// 		DepartmentID: to.Ptr("164821"),
	// 		DisplayName: to.Ptr("AccountName0"),
	// 		EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-31T17:32:28.000Z"); return t}()),
	// 		IsDevTestEnabled: to.Ptr(true),
	// 		StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
	// 		Status: to.Ptr("Active"),
	// 	},
	// }
}
