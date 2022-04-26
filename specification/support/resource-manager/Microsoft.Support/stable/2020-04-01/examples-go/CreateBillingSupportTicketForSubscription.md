Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv0.4.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateBillingSupportTicketForSubscription.json
func ExampleTicketsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsupport.NewTicketsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<support-ticket-name>",
		armsupport.TicketDetails{
			Properties: &armsupport.TicketDetailsProperties{
				Description: to.Ptr("<description>"),
				ContactDetails: &armsupport.ContactProfile{
					Country:                  to.Ptr("<country>"),
					FirstName:                to.Ptr("<first-name>"),
					LastName:                 to.Ptr("<last-name>"),
					PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
					PreferredSupportLanguage: to.Ptr("<preferred-support-language>"),
					PreferredTimeZone:        to.Ptr("<preferred-time-zone>"),
					PrimaryEmailAddress:      to.Ptr("<primary-email-address>"),
				},
				ProblemClassificationID: to.Ptr("<problem-classification-id>"),
				ServiceID:               to.Ptr("<service-id>"),
				Severity:                to.Ptr(armsupport.SeverityLevelModerate),
				Title:                   to.Ptr("<title>"),
			},
		},
		&armsupport.TicketsClientBeginCreateOptions{ResumeToken: ""})
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
