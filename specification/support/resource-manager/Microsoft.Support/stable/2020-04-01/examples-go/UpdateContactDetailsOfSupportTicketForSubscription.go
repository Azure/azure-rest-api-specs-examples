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
	}
	ctx := context.Background()
	client, err := armsupport.NewTicketsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"testticket",
		armsupport.UpdateSupportTicket{
			ContactDetails: &armsupport.UpdateContactProfile{
				AdditionalEmailAddresses: []*string{
					to.Ptr("tname@contoso.com"),
					to.Ptr("teamtest@contoso.com")},
				Country:                  to.Ptr("USA"),
				FirstName:                to.Ptr("first name"),
				LastName:                 to.Ptr("last name"),
				PhoneNumber:              to.Ptr("123-456-7890"),
				PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
				PreferredSupportLanguage: to.Ptr("en-US"),
				PreferredTimeZone:        to.Ptr("Pacific Standard Time"),
				PrimaryEmailAddress:      to.Ptr("test.name@contoso.com"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
