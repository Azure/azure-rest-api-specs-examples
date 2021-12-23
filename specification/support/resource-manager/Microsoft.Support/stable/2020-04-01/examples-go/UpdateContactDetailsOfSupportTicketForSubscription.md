Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv0.1.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.

```go
package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/UpdateContactDetailsOfSupportTicketForSubscription.json
func ExampleSupportTicketsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewSupportTicketsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<support-ticket-name>",
		armsupport.UpdateSupportTicket{
			ContactDetails: &armsupport.UpdateContactProfile{
				AdditionalEmailAddresses: []*string{
					to.StringPtr("tname@contoso.com"),
					to.StringPtr("teamtest@contoso.com")},
				Country:                  to.StringPtr("<country>"),
				FirstName:                to.StringPtr("<first-name>"),
				LastName:                 to.StringPtr("<last-name>"),
				PhoneNumber:              to.StringPtr("<phone-number>"),
				PreferredContactMethod:   armsupport.PreferredContactMethodEmail.ToPtr(),
				PreferredSupportLanguage: to.StringPtr("<preferred-support-language>"),
				PreferredTimeZone:        to.StringPtr("<preferred-time-zone>"),
				PrimaryEmailAddress:      to.StringPtr("<primary-email-address>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SupportTicketDetails.ID: %s\n", *res.ID)
}
```
