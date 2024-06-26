package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateReservationDetailsReportByBillingAccount.json
func ExampleGenerateReservationDetailsReportClient_BeginByBillingAccountID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGenerateReservationDetailsReportClient().BeginByBillingAccountID(ctx, "9845612", "2020-01-01", "2020-01-30", nil)
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
	// res.OperationStatus = armcostmanagement.OperationStatus{
	// 	Properties: &armcostmanagement.ReportURL{
	// 		ReportURL: to.Ptr(armcostmanagement.ReservationReportSchema("https://storage.blob.core.windows.net/details/20200911/00000000-0000-0000-0000-000000000000?sv=2016-05-31&sr=b&sig=jep8HT2aphfUkyERRZa5LRfd9RPzjXbzB%2F9TNiQ")),
	// 		ValidUntil: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-12T02:56:55.502Z"); return t}()),
	// 	},
	// 	Status: to.Ptr(armcostmanagement.OperationStatusTypeCompleted),
	// }
}
