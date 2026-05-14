package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/departmentsListByBillingAccount.json
func ExampleDepartmentsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDepartmentsClient().NewListByBillingAccountPager("456598", nil)
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
		// page = armbilling.DepartmentsClientListByBillingAccountResponse{
		// 	DepartmentListResult: armbilling.DepartmentListResult{
		// 		Value: []*armbilling.Department{
		// 			{
		// 				Name: to.Ptr("164821"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/departments"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/456598/departments/164821"),
		// 				Properties: &armbilling.DepartmentProperties{
		// 					CostCenter: to.Ptr("C1"),
		// 					DisplayName: to.Ptr("departmentName1"),
		// 					Status: to.Ptr("Active"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("164822"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/departments"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/456598/departments/164822"),
		// 				Properties: &armbilling.DepartmentProperties{
		// 					CostCenter: to.Ptr("C4"),
		// 					DisplayName: to.Ptr("departmentName2"),
		// 					Status: to.Ptr("Active"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
