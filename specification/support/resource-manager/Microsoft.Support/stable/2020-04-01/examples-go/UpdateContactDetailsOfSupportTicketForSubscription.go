package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/UpdateContactDetailsOfSupportTicketForSubscription.json
func ExampleTicketsClient_Update_updateContactDetailsOfASupportTicket() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTicketsClient().Update(ctx, "testticket", armsupport.UpdateSupportTicket{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TicketDetails = armsupport.TicketDetails{
	// 	Name: to.Ptr("testticket"),
	// 	Type: to.Ptr("Microsoft.Support/supportTickets"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket"),
	// 	Properties: &armsupport.TicketDetailsProperties{
	// 		Description: to.Ptr("This is a test - please ignore"),
	// 		ContactDetails: &armsupport.ContactProfile{
	// 			AdditionalEmailAddresses: []*string{
	// 				to.Ptr("tname@contoso.com"),
	// 				to.Ptr("teamtest@contoso.com")},
	// 				Country: to.Ptr("USA"),
	// 				FirstName: to.Ptr("first name"),
	// 				LastName: to.Ptr("last name"),
	// 				PhoneNumber: to.Ptr("123-456-7890"),
	// 				PreferredContactMethod: to.Ptr(armsupport.PreferredContactMethodEmail),
	// 				PreferredSupportLanguage: to.Ptr("en-US"),
	// 				PreferredTimeZone: to.Ptr("Pacific Standard Time"),
	// 				PrimaryEmailAddress: to.Ptr("test.name@contoso.com"),
	// 			},
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
	// 			ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:23.000Z"); return t}()),
	// 			ProblemClassificationDisplayName: to.Ptr("Add or Edit VAT, TAX ID, or PO Number"),
	// 			ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid/problemClassifications/problemClassification_guid"),
	// 			Require24X7Response: to.Ptr(false),
	// 			ServiceDisplayName: to.Ptr("Subscription management"),
	// 			ServiceID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid"),
	// 			ServiceLevelAgreement: &armsupport.ServiceLevelAgreement{
	// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-21T17:36:18.000Z"); return t}()),
	// 				SLAMinutes: to.Ptr[int32](240),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
	// 			},
	// 			Severity: to.Ptr(armsupport.SeverityLevelCritical),
	// 			Status: to.Ptr("Open"),
	// 			SupportEngineer: &armsupport.Engineer{
	// 				EmailAddress: to.Ptr("xyz@contoso.com"),
	// 			},
	// 			SupportPlanType: to.Ptr("Premier"),
	// 			SupportTicketID: to.Ptr("118032014183770"),
	// 			Title: to.Ptr("Test - please ignore"),
	// 		},
	// 	}
}
