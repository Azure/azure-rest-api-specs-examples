package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListWebCommunicationsForSupportTicket.json
func ExampleCommunicationsNoSubscriptionClient_NewListPager_listWebCommunicationsForANoSubscriptionSupportTicket() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunicationsNoSubscriptionClient().NewListPager("testticket", &armsupport.CommunicationsNoSubscriptionClientListOptions{Top: nil,
		Filter: to.Ptr("communicationType eq 'web'"),
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
		// page.CommunicationsListResult = armsupport.CommunicationsListResult{
		// 	Value: []*armsupport.CommunicationDetails{
		// 		{
		// 			Name: to.Ptr("testmessage1"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage1"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("this is a test message"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("this is a test message"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testmessage2"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage2"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("test"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T10:53:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("test"),
		// 			},
		// 	}},
		// }
	}
}
