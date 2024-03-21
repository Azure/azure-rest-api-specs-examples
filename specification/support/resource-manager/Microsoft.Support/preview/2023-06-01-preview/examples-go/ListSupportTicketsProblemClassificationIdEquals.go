package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ListSupportTicketsProblemClassificationIdEquals.json
func ExampleTicketsNoSubscriptionClient_NewListPager_listSupportTicketsWithACertainProblemClassificationId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTicketsNoSubscriptionClient().NewListPager(&armsupport.TicketsNoSubscriptionClientListOptions{Top: nil,
		Filter: to.Ptr("ProblemClassificationId eq 'compute_vm_problemClassification_guid'"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TicketsListResult = armsupport.TicketsListResult{
		// 	Value: []*armsupport.TicketDetails{
		// 		{
		// 			Name: to.Ptr("testTicket1"),
		// 			Type: to.Ptr("Microsoft.Support/supportTickets"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testTicket1"),
		// 			Properties: &armsupport.TicketDetailsProperties{
		// 				Description: to.Ptr("my description"),
		// 				AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentYes),
		// 				ContactDetails: &armsupport.ContactProfile{
		// 					Country: to.Ptr("usa"),
		// 					FirstName: to.Ptr("abc"),
		// 					LastName: to.Ptr("xyz"),
		// 					PreferredContactMethod: to.Ptr(armsupport.PreferredContactMethodEmail),
		// 					PreferredSupportLanguage: to.Ptr("en-US"),
		// 					PreferredTimeZone: to.Ptr("Pacific Standard Time"),
		// 					PrimaryEmailAddress: to.Ptr("abc@contoso.com"),
		// 				},
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T21:52:10.000Z"); return t}()),
		// 				EnrollmentID: to.Ptr(""),
		// 				FileWorkspaceName: to.Ptr("testTicket1"),
		// 				ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-12T23:05:19.000Z"); return t}()),
		// 				ProblemClassificationDisplayName: to.Ptr("Compute-VM (cores-vCPUs) subscription limit increases"),
		// 				ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/compute_vm_problemClassification_guid"),
		// 				Require24X7Response: to.Ptr(false),
		// 				ServiceDisplayName: to.Ptr("service_displayName"),
		// 				ServiceID: to.Ptr("/providers/Microsoft.Support/services/service_guid"),
		// 				Severity: to.Ptr(armsupport.SeverityLevelMinimal),
		// 				Status: to.Ptr("Open"),
		// 				SupportEngineer: &armsupport.Engineer{
		// 					EmailAddress: to.Ptr("xyz@contoso.com"),
		// 				},
		// 				SupportPlanDisplayName: to.Ptr("Premier"),
		// 				SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
		// 				SupportPlanType: to.Ptr("Premier"),
		// 				SupportTicketID: to.Ptr("2205060010000072"),
		// 				Title: to.Ptr("my title"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testTicket2"),
		// 			Type: to.Ptr("Microsoft.Support/supportTickets"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testTicket2"),
		// 			Properties: &armsupport.TicketDetailsProperties{
		// 				Description: to.Ptr("This is a test - please ignore"),
		// 				AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentYes),
		// 				ContactDetails: &armsupport.ContactProfile{
		// 					Country: to.Ptr("USA"),
		// 					FirstName: to.Ptr("abc"),
		// 					LastName: to.Ptr("xyz"),
		// 					PreferredContactMethod: to.Ptr(armsupport.PreferredContactMethodEmail),
		// 					PreferredSupportLanguage: to.Ptr("en-US"),
		// 					PreferredTimeZone: to.Ptr("Pacific Standard Time"),
		// 					PrimaryEmailAddress: to.Ptr("abc@contoso.com"),
		// 				},
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T21:38:42.000Z"); return t}()),
		// 				FileWorkspaceName: to.Ptr("testTicket2"),
		// 				ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T21:39:14.000Z"); return t}()),
		// 				ProblemClassificationDisplayName: to.Ptr("Compute-VM (cores-vCPUs) subscription limit increases"),
		// 				ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/compute_vm_problemClassification_guid"),
		// 				Require24X7Response: to.Ptr(false),
		// 				ServiceDisplayName: to.Ptr("service_displayName"),
		// 				ServiceID: to.Ptr("/providers/Microsoft.Support/services/service_guid"),
		// 				Severity: to.Ptr(armsupport.SeverityLevelMinimal),
		// 				Status: to.Ptr("Open"),
		// 				SupportEngineer: &armsupport.Engineer{
		// 					EmailAddress: to.Ptr("xyz@contoso.com"),
		// 				},
		// 				SupportPlanDisplayName: to.Ptr("Premier"),
		// 				SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
		// 				SupportPlanType: to.Ptr("Premier"),
		// 				SupportTicketID: to.Ptr("2205040010000077"),
		// 				Title: to.Ptr("Test - please ignore"),
		// 			},
		// 	}},
		// }
	}
}
