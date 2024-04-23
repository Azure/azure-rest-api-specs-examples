package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListSupportTicketsInUpdatingState.json
func ExampleTicketsNoSubscriptionClient_NewListPager_listSupportTicketsInUpdatingState() {
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
		Filter: to.Ptr("status eq 'Updating'"),
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
		// 			Name: to.Ptr("testticket"),
		// 			Type: to.Ptr("Microsoft.Support/supportTickets"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket"),
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
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
		// 				FileWorkspaceName: to.Ptr("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066"),
		// 				ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:23.000Z"); return t}()),
		// 				ProblemClassificationDisplayName: to.Ptr("Connectivity / Cannot connect to virtual machine by using RDP or SSH"),
		// 				ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/virtual_machine_service_guid/problemClassifications/problemClassification_guid"),
		// 				ProblemScopingQuestions: to.Ptr("{\"articleId\":\"076846c1-4c0b-4b21-91c6-1a30246b3867\",\"scopingDetails\":[{\"question\":\"When did the problem begin?\",\"controlId\":\"problem_start_time\",\"orderId\":1,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"2023-08-31T18:55:00.739Z\",\"value\":\"2023-08-31T18:55:00.739Z\",\"type\":\"datetime\"}},{\"question\":\"API Type of the Cosmos DB account\",\"controlId\":\"api_type\",\"orderId\":2,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"Table\",\"value\":\"tables\",\"type\":\"string\"}},{\"question\":\"Table name\",\"controlId\":\"collection_name_table\",\"orderId\":11,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"Select Table Name\",\"value\":\"dont_know_answer\",\"type\":\"string\"}},{\"question\":\"Provide additional details about the issue you're facing\",\"controlId\":\"problem_description\",\"orderId\":12,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"test ticket, please ignore and close\",\"value\":\"test ticket, please ignore and close\",\"type\":\"string\"}}]}"),
		// 				Require24X7Response: to.Ptr(false),
		// 				SecondaryConsent: []*armsupport.SecondaryConsent{
		// 					{
		// 						Type: to.Ptr("VirtualMachine"),
		// 						UserConsent: to.Ptr(armsupport.UserConsentYes),
		// 				}},
		// 				ServiceDisplayName: to.Ptr("Virtual Machine running Linux"),
		// 				ServiceID: to.Ptr("/providers/Microsoft.Support/services/virtual_machine_service_guid"),
		// 				ServiceLevelAgreement: &armsupport.ServiceLevelAgreement{
		// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-21T17:36:18.000Z"); return t}()),
		// 					SLAMinutes: to.Ptr[int32](240),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
		// 				},
		// 				Severity: to.Ptr(armsupport.SeverityLevelModerate),
		// 				Status: to.Ptr("Updating"),
		// 				SupportEngineer: &armsupport.Engineer{
		// 					EmailAddress: to.Ptr("xyz@contoso.com"),
		// 				},
		// 				SupportPlanDisplayName: to.Ptr("Premier"),
		// 				SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
		// 				SupportPlanType: to.Ptr("Premier"),
		// 				SupportTicketID: to.Ptr("119120321001170"),
		// 				Title: to.Ptr("my title"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testticket2"),
		// 			Type: to.Ptr("Microsoft.Support/supportTickets"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket2"),
		// 			Properties: &armsupport.TicketDetailsProperties{
		// 				Description: to.Ptr("This is a test - please ignore"),
		// 				AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentNo),
		// 				ContactDetails: &armsupport.ContactProfile{
		// 					Country: to.Ptr("USA"),
		// 					FirstName: to.Ptr("abc"),
		// 					LastName: to.Ptr("xyz"),
		// 					PreferredContactMethod: to.Ptr(armsupport.PreferredContactMethodEmail),
		// 					PreferredSupportLanguage: to.Ptr("en-US"),
		// 					PreferredTimeZone: to.Ptr("Pacific Standard Time"),
		// 					PrimaryEmailAddress: to.Ptr("abc@contoso.com"),
		// 				},
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
		// 				FileWorkspaceName: to.Ptr("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066"),
		// 				ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:23.000Z"); return t}()),
		// 				ProblemClassificationDisplayName: to.Ptr("Add or Edit VAT, TAX ID, or PO Number"),
		// 				ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid/problemClassifications/problemClassification_guid"),
		// 				Require24X7Response: to.Ptr(false),
		// 				ServiceDisplayName: to.Ptr("Subscription management"),
		// 				ServiceID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid"),
		// 				ServiceLevelAgreement: &armsupport.ServiceLevelAgreement{
		// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-21T17:36:18.000Z"); return t}()),
		// 					SLAMinutes: to.Ptr[int32](240),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
		// 				},
		// 				Severity: to.Ptr(armsupport.SeverityLevelMinimal),
		// 				Status: to.Ptr("Updating"),
		// 				SupportEngineer: &armsupport.Engineer{
		// 					EmailAddress: to.Ptr("xyz@contoso.com"),
		// 				},
		// 				SupportPlanDisplayName: to.Ptr("Premier"),
		// 				SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
		// 				SupportPlanType: to.Ptr("Premier"),
		// 				SupportTicketID: to.Ptr("118032014183771"),
		// 				Title: to.Ptr("Test - please ignore"),
		// 			},
		// 	}},
		// }
	}
}
