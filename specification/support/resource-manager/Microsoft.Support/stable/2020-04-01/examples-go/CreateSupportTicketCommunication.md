Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv0.1.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.

```go
package armsupport_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateSupportTicketCommunication.json
func ExampleCommunicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewCommunicationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<support-ticket-name>",
		"<communication-name>",
		armsupport.CommunicationDetails{
			Properties: &armsupport.CommunicationDetailsProperties{
				Body:    to.StringPtr("<body>"),
				Sender:  to.StringPtr("<sender>"),
				Subject: to.StringPtr("<subject>"),
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
	log.Printf("CommunicationDetails.ID: %s\n", *res.ID)
}
```
