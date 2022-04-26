Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv0.4.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.

```go
package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/UpdateContactDetailsOfSupportTicketForSubscription.json
func ExampleTicketsClient_Update() {
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
	res, err := client.Update(ctx,
		"<support-ticket-name>",
		armsupport.UpdateSupportTicket{
			ContactDetails: &armsupport.UpdateContactProfile{
				AdditionalEmailAddresses: []*string{
					to.Ptr("tname@contoso.com"),
					to.Ptr("teamtest@contoso.com")},
				Country:                  to.Ptr("<country>"),
				FirstName:                to.Ptr("<first-name>"),
				LastName:                 to.Ptr("<last-name>"),
				PhoneNumber:              to.Ptr("<phone-number>"),
				PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
				PreferredSupportLanguage: to.Ptr("<preferred-support-language>"),
				PreferredTimeZone:        to.Ptr("<preferred-time-zone>"),
				PrimaryEmailAddress:      to.Ptr("<primary-email-address>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
