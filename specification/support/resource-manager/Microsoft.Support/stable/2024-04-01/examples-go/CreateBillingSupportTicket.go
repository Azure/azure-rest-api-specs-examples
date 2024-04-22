package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateBillingSupportTicket.json
func ExampleTicketsNoSubscriptionClient_BeginCreate_createATicketForBillingRelatedIssues() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTicketsNoSubscriptionClient().BeginCreate(ctx, "testticket", armsupport.TicketDetails{
		Properties: &armsupport.TicketDetailsProperties{
			Description:               to.Ptr("my description"),
			AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentNo),
			ContactDetails: &armsupport.ContactProfile{
				Country:                  to.Ptr("usa"),
				FirstName:                to.Ptr("abc"),
				LastName:                 to.Ptr("xyz"),
				PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
				PreferredSupportLanguage: to.Ptr("en-US"),
				PreferredTimeZone:        to.Ptr("Pacific Standard Time"),
				PrimaryEmailAddress:      to.Ptr("abc@contoso.com"),
			},
			FileWorkspaceName:       to.Ptr("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066"),
			ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid"),
			ServiceID:               to.Ptr("/providers/Microsoft.Support/services/billing_service_guid"),
			Severity:                to.Ptr(armsupport.SeverityLevelModerate),
			SupportPlanID:           to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
			Title:                   to.Ptr("my title"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TicketDetails = armsupport.TicketDetails{
	// 	Name: to.Ptr("testticket"),
	// 	Type: to.Ptr("Microsoft.Support/supportTickets"),
	// 	ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket"),
	// 	Properties: &armsupport.TicketDetailsProperties{
	// 		Description: to.Ptr("my description"),
	// 		AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentNo),
	// 		ContactDetails: &armsupport.ContactProfile{
	// 			Country: to.Ptr("usa"),
	// 			FirstName: to.Ptr("abc"),
	// 			LastName: to.Ptr("xyz"),
	// 			PreferredContactMethod: to.Ptr(armsupport.PreferredContactMethodEmail),
	// 			PreferredSupportLanguage: to.Ptr("en-US"),
	// 			PreferredTimeZone: to.Ptr("Pacific Standard Time"),
	// 			PrimaryEmailAddress: to.Ptr("abc@contoso.com"),
	// 		},
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
	// 		FileWorkspaceName: to.Ptr("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066"),
	// 		ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:23.000Z"); return t}()),
	// 		ProblemClassificationDisplayName: to.Ptr("Refund request"),
	// 		ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid"),
	// 		Require24X7Response: to.Ptr(false),
	// 		ServiceDisplayName: to.Ptr("Billing"),
	// 		ServiceID: to.Ptr("/providers/Microsoft.Support/services/billing_service_guid"),
	// 		ServiceLevelAgreement: &armsupport.ServiceLevelAgreement{
	// 			ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-21T17:36:18.000Z"); return t}()),
	// 			SLAMinutes: to.Ptr[int32](240),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
	// 		},
	// 		Severity: to.Ptr(armsupport.SeverityLevelModerate),
	// 		Status: to.Ptr("Open"),
	// 		SupportEngineer: &armsupport.Engineer{
	// 		},
	// 		SupportPlanDisplayName: to.Ptr("Premier"),
	// 		SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
	// 		SupportPlanType: to.Ptr("Premier"),
	// 		SupportTicketID: to.Ptr("119120321001170"),
	// 		Title: to.Ptr("my title"),
	// 	},
	// }
}
