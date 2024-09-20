package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/departmentGet.json
func ExampleDepartmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDepartmentsClient().Get(ctx, "456598", "164821", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Department = armbilling.Department{
	// 	Name: to.Ptr("164821"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/departments"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/456598/departments/164821"),
	// 	Properties: &armbilling.DepartmentProperties{
	// 		CostCenter: to.Ptr("C1"),
	// 		DisplayName: to.Ptr("Test department"),
	// 		Status: to.Ptr("Active"),
	// 	},
	// }
}
