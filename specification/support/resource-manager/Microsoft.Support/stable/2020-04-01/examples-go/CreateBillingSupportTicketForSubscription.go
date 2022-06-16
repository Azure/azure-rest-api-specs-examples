package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateBillingSupportTicketForSubscription.json
func ExampleTicketsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsupport.NewTicketsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testticket",
		armsupport.TicketDetails{
			Properties: &armsupport.TicketDetailsProperties{
				Description: to.Ptr("my description"),
				ContactDetails: &armsupport.ContactProfile{
					Country:                  to.Ptr("usa"),
					FirstName:                to.Ptr("abc"),
					LastName:                 to.Ptr("xyz"),
					PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
					PreferredSupportLanguage: to.Ptr("en-US"),
					PreferredTimeZone:        to.Ptr("Pacific Standard Time"),
					PrimaryEmailAddress:      to.Ptr("abc@contoso.com"),
				},
				ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid"),
				ServiceID:               to.Ptr("/providers/Microsoft.Support/services/billing_service_guid"),
				Severity:                to.Ptr(armsupport.SeverityLevelModerate),
				Title:                   to.Ptr("my title"),
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
