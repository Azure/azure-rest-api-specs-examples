Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.4.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/CustomerEventCreate.json
func ExampleCustomerEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtestbase.NewCustomerEventsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		"<customer-event-name>",
		armtestbase.CustomerEventResource{
			Properties: &armtestbase.CustomerEventProperties{
				EventName: to.Ptr("<event-name>"),
				Receivers: []*armtestbase.NotificationEventReceiver{
					{
						ReceiverType: to.Ptr("<receiver-type>"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							UserObjectReceiverValue: &armtestbase.UserObjectReceiverValue{
								UserObjectIDs: []*string{
									to.Ptr("245245245245325"),
									to.Ptr("365365365363565")},
							},
						},
					},
					{
						ReceiverType: to.Ptr("<receiver-type>"),
						ReceiverValue: &armtestbase.NotificationReceiverValue{
							DistributionGroupListReceiverValue: &armtestbase.DistributionGroupListReceiverValue{
								DistributionGroups: []*string{
									to.Ptr("test@microsoft.com")},
							},
						},
					}},
			},
		},
		&armtestbase.CustomerEventsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
