package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/EAPriceSheetForBillingPeriod.json
func ExamplePriceSheetClient_BeginDownloadByBillingAccountPeriod() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPriceSheetClient().BeginDownloadByBillingAccountPeriod(ctx, "0000000", "202305", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.PriceSheetClientDownloadByBillingAccountPeriodResponse{
	// 	OperationStatus: armconsumption.OperationStatus{
	// 		Properties: &armconsumption.PricesheetDownloadProperties{
	// 			DownloadURL: to.Ptr("https://xxxxxx.blob.core.windows.net/armpricesheetreportdownloadcontainer/20230510/00000000-0000-0000-0000-000000000000"),
	// 			ValidTill: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T05:12:50.4266333Z"); return t}()),
	// 		},
	// 		Status: to.Ptr(armconsumption.OperationStatusTypeCompleted),
	// 	},
	// }
}
