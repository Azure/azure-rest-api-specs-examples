Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.2.1/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/CustomerEventCreate.json
func ExampleCustomerEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewCustomerEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<customer-event-name>",
		armtestbase.CustomerEventResource{
			Properties: &armtestbase.CustomerEventProperties{
				EventName: to.StringPtr("<event-name>"),
				Receivers: []*armtestbase.NotificationEventReceiver{
					{
						ReceiverType: to.StringPtr("<receiver-type>"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							UserObjectReceiverValue: &armtestbase.UserObjectReceiverValue{
								UserObjectIDs: []*string{
									to.StringPtr("245245245245325"),
									to.StringPtr("365365365363565")},
							},
						},
					},
					{
						ReceiverType: to.StringPtr("<receiver-type>"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							DistributionGroupListReceiverValue: &armtestbase.DistributionGroupListReceiverValue{
								DistributionGroups: []*string{
									to.StringPtr("test@microsoft.com")},
							},
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CustomerEventsClientCreateResult)
}
```
