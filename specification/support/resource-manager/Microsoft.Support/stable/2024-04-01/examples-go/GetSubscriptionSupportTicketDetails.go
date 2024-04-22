package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetSubscriptionSupportTicketDetails.json
func ExampleTicketsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTicketsClient().Get(ctx, "testticket", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TicketDetails = armsupport.TicketDetails{
	// 	Name: to.Ptr("testticket"),
	// 	Type: to.Ptr("Microsoft.Support/supportTickets"),
	// 	ID: to.Ptr("/subscriptions/132d901f-189d-4381-9214-fe68e27e05a1/providers/Microsoft.Support/supportTickets/testticket"),
	// 	Properties: &armsupport.TicketDetailsProperties{
	// 		Description: to.Ptr("This is a test - please ignore"),
	// 		AdvancedDiagnosticConsent: to.Ptr(armsupport.ConsentNo),
	// 		ContactDetails: &armsupport.ContactProfile{
	// 			Country: to.Ptr("USA"),
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
	// 		ProblemClassificationDisplayName: to.Ptr("Add or Edit VAT, TAX ID, or PO Number"),
	// 		ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid/problemClassifications/problemClassification_guid"),
	// 		Require24X7Response: to.Ptr(false),
	// 		ServiceDisplayName: to.Ptr("Subscription management"),
	// 		ServiceID: to.Ptr("/providers/Microsoft.Support/services/subscription_management_service_guid"),
	// 		ServiceLevelAgreement: &armsupport.ServiceLevelAgreement{
	// 			ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-21T17:36:18.000Z"); return t}()),
	// 			SLAMinutes: to.Ptr[int32](240),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:18.000Z"); return t}()),
	// 		},
	// 		Severity: to.Ptr(armsupport.SeverityLevelMinimal),
	// 		Status: to.Ptr("Open"),
	// 		SupportEngineer: &armsupport.Engineer{
	// 			EmailAddress: to.Ptr("xyz@contoso.com"),
	// 		},
	// 		SupportPlanDisplayName: to.Ptr("Premier"),
	// 		SupportPlanID: to.Ptr("U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw="),
	// 		SupportPlanType: to.Ptr("Premier"),
	// 		SupportTicketID: to.Ptr("118032014183770"),
	// 		Title: to.Ptr("Test - please ignore"),
	// 	},
	// }
}
