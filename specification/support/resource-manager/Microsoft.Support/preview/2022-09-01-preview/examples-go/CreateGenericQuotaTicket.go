package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateGenericQuotaTicket.json
func ExampleTicketsClient_BeginCreate_createATicketToRequestQuotaIncreaseForServicesThatDoNotRequireAdditionalDetailsInTheQuotaTicketDetailsObject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTicketsClient().BeginCreate(ctx, "testticket", armsupport.TicketDetails{
		Properties: &armsupport.TicketDetailsProperties{
			Description: to.Ptr("Increase the maximum throughput per container limit to 10000 for account foo bar"),
			ContactDetails: &armsupport.ContactProfile{
				Country:                  to.Ptr("usa"),
				FirstName:                to.Ptr("abc"),
				LastName:                 to.Ptr("xyz"),
				PreferredContactMethod:   to.Ptr(armsupport.PreferredContactMethodEmail),
				PreferredSupportLanguage: to.Ptr("en-US"),
				PreferredTimeZone:        to.Ptr("Pacific Standard Time"),
				PrimaryEmailAddress:      to.Ptr("abc@contoso.com"),
			},
			ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid"),
			ServiceID:               to.Ptr("/providers/Microsoft.Support/services/quota_service_guid"),
			Severity:                to.Ptr(armsupport.SeverityLevelModerate),
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
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket"),
	// 	Properties: &armsupport.TicketDetailsProperties{
	// 		Description: to.Ptr("Increase the maximum throughput per container limit to 10000 for account foo bar"),
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
	// 		ModifiedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-20T21:36:23.000Z"); return t}()),
	// 		ProblemClassificationDisplayName: to.Ptr("Cosmos DB"),
	// 		ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid"),
	// 		Require24X7Response: to.Ptr(false),
	// 		ServiceDisplayName: to.Ptr("Service and subscription limits (quotas)"),
	// 		ServiceID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid"),
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
	// 		SupportPlanType: to.Ptr("Premier"),
	// 		SupportTicketID: to.Ptr("119120321001170"),
	// 		Title: to.Ptr("my title"),
	// 	},
	// }
}
