package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateCoresQuotaTicketForSubscription.json
func ExampleTicketsClient_BeginCreate_createATicketToRequestQuotaIncreaseForComputeVmCores() {
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
			ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cores_problemClassification_guid"),
			QuotaTicketDetails: &armsupport.QuotaTicketDetails{
				QuotaChangeRequestVersion: to.Ptr("1.0"),
				QuotaChangeRequests: []*armsupport.QuotaChangeRequest{
					{
						Payload: to.Ptr("{\"SKU\":\"DSv3 Series\",\"NewLimit\":104}"),
						Region:  to.Ptr("EastUS"),
					}},
			},
			ServiceID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid"),
			Severity:  to.Ptr(armsupport.SeverityLevelModerate),
			Title:     to.Ptr("my title"),
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
	// 		Description: to.Ptr("my description"),
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
	// 		ProblemClassificationDisplayName: to.Ptr("Compute-VM (cores-vCPUs) subscription limit increases"),
	// 		ProblemClassificationID: to.Ptr("/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cores_problemClassification_guid"),
	// 		QuotaTicketDetails: &armsupport.QuotaTicketDetails{
	// 			QuotaChangeRequestVersion: to.Ptr("1.0"),
	// 			QuotaChangeRequests: []*armsupport.QuotaChangeRequest{
	// 				{
	// 					Payload: to.Ptr("{\"VmFamily\":\"DSv3 Series\",\"NewLimit\":104}"),
	// 					Region: to.Ptr("EastUS"),
	// 			}},
	// 		},
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
	// 		SupportPlanType: to.Ptr("Premier"),
	// 		SupportTicketID: to.Ptr("119120321001170"),
	// 		Title: to.Ptr("my title"),
	// 	},
	// }
}
