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

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateBillingSupportTicketForSubscription.json
func ExampleSupportTicketsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewSupportTicketsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<support-ticket-name>",
		armsupport.SupportTicketDetails{
			Properties: &armsupport.SupportTicketDetailsProperties{
				Description: to.StringPtr("<description>"),
				ContactDetails: &armsupport.ContactProfile{
					Country:                  to.StringPtr("<country>"),
					FirstName:                to.StringPtr("<first-name>"),
					LastName:                 to.StringPtr("<last-name>"),
					PreferredContactMethod:   armsupport.PreferredContactMethodEmail.ToPtr(),
					PreferredSupportLanguage: to.StringPtr("<preferred-support-language>"),
					PreferredTimeZone:        to.StringPtr("<preferred-time-zone>"),
					PrimaryEmailAddress:      to.StringPtr("<primary-email-address>"),
				},
				ProblemClassificationID: to.StringPtr("<problem-classification-id>"),
				ServiceID:               to.StringPtr("<service-id>"),
				Severity:                armsupport.SeverityLevelModerate.ToPtr(),
				Title:                   to.StringPtr("<title>"),
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
	log.Printf("SupportTicketDetails.ID: %s\n", *res.ID)
}
```
