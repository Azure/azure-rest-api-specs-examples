package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/GenerateDetailedCostReportByBillingAccountLegacyAndBillingPeriod.json
func ExampleGenerateDetailedCostReportClient_BeginCreateOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewGenerateDetailedCostReportClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOperation(ctx,
		"providers/Microsoft.Billing/billingAccounts/12345",
		armcostmanagement.GenerateDetailedCostReportDefinition{
			BillingPeriod: to.Ptr("202008"),
			Metric:        to.Ptr(armcostmanagement.GenerateDetailedCostReportMetricTypeActualCost),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
