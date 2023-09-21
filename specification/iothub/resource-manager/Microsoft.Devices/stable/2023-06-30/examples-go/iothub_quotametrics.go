package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_quotametrics.json
func ExampleResourceClient_NewGetQuotaMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceClient().NewGetQuotaMetricsPager("myResourceGroup", "testHub", nil)
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
		// page.QuotaMetricInfoListResult = armiothub.QuotaMetricInfoListResult{
		// 	Value: []*armiothub.QuotaMetricInfo{
		// 		{
		// 			Name: to.Ptr("TotalMessages"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			MaxValue: to.Ptr[int64](400000),
		// 		},
		// 		{
		// 			Name: to.Ptr("TotalDeviceCount"),
		// 			CurrentValue: to.Ptr[int64](0),
		// 			MaxValue: to.Ptr[int64](500000),
		// 	}},
		// }
	}
}
