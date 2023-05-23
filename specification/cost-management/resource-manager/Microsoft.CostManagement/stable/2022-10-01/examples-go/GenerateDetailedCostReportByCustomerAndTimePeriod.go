package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportByCustomerAndTimePeriod.json
func ExampleGenerateDetailedCostReportClient_BeginCreateOperation_generateDetailedCostReportByCustomerAndTimePeriod() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGenerateDetailedCostReportClient().BeginCreateOperation(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/13579", armcostmanagement.GenerateDetailedCostReportDefinition{
		Metric: to.Ptr(armcostmanagement.GenerateDetailedCostReportMetricTypeActualCost),
		TimePeriod: &armcostmanagement.GenerateDetailedCostReportTimePeriod{
			End:   to.Ptr("2020-03-15"),
			Start: to.Ptr("2020-03-01"),
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
	// res.GenerateDetailedCostReportOperationResult = armcostmanagement.GenerateDetailedCostReportOperationResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Consumption/operationResult"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/operationResults/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.DownloadURL{
	// 		DownloadURL: to.Ptr("https://ccmreportstorageeastus.blob.core.windows.net/armreports/20201207/00000000-0000-0000-0000-000000000000?sv=2020-05-31&sr=b&sig=abcd"),
	// 		ValidTill: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-08T05:55:59.4394737Z"); return t}()),
	// 	},
	// }
}
