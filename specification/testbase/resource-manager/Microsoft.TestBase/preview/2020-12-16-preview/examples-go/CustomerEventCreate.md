```go
package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/CustomerEventCreate.json
func ExampleCustomerEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtestbase.NewCustomerEventsClient("476f61a4-952c-422a-b4db-568a828f35df", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contoso-rg1",
		"contoso-testBaseAccount1",
		"WeeklySummary",
		armtestbase.CustomerEventResource{
			Properties: &armtestbase.CustomerEventProperties{
				EventName: to.Ptr("WeeklySummary"),
				Receivers: []*armtestbase.NotificationEventReceiver{
					{
						ReceiverType: to.Ptr("UserObjects"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							UserObjectReceiverValue: &armtestbase.UserObjectReceiverValue{
								UserObjectIDs: []*string{
									to.Ptr("245245245245325"),
									to.Ptr("365365365363565")},
							},
						},
					},
					{
						ReceiverType: to.Ptr("DistributionGroup"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							DistributionGroupListReceiverValue: &armtestbase.DistributionGroupListReceiverValue{
								DistributionGroups: []*string{
									to.Ptr("test@microsoft.com")},
							},
						},
					}},
			},
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.5.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.
