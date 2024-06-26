package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListChatTranscriptsForSubscriptionSupportTicket.json
func ExampleChatTranscriptsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChatTranscriptsClient().NewListPager("testticket", nil)
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
		// page.ChatTranscriptsListResult = armsupport.ChatTranscriptsListResult{
		// 	Value: []*armsupport.ChatTranscriptDetails{
		// 		{
		// 			Name: to.Ptr("55989c71-1727-4cd9-abad-ddb8770f71cd"),
		// 			Type: to.Ptr("Microsoft.Support/chatTranscripts"),
		// 			ID: to.Ptr("/subscriptions/132d901f-189d-4381-9214-fe68e27e05a1/providers/Microsoft.Support/supportTickets/2207120020000085/chatTranscripts/55989c71-1727-4cd9-abad-ddb8770f71cd"),
		// 			Properties: &armsupport.ChatTranscriptDetailsProperties{
		// 				Messages: []*armsupport.MessageProperties{
		// 					{
		// 						Body: to.Ptr("Hi"),
		// 						CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 						ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
		// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-24T20:18:19.000Z"); return t}()),
		// 						Sender: to.Ptr("support engineer"),
		// 					},
		// 					{
		// 						Body: to.Ptr("hi"),
		// 						CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionInbound),
		// 						ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
		// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-24T20:19:16.000Z"); return t}()),
		// 						Sender: to.Ptr("user"),
		// 				}},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-22T22:46:35.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("f15051e3-a2f2-489f-9e64-8cfa203f44f8"),
		// 			Type: to.Ptr("Microsoft.Support/chatTranscripts"),
		// 			ID: to.Ptr("/subscriptions/132d901f-189d-4381-9214-fe68e27e05a1/providers/Microsoft.Support/supportTickets/2207120020000085/chatTranscripts/f15051e3-a2f2-489f-9e64-8cfa203f44f8"),
		// 			Properties: &armsupport.ChatTranscriptDetailsProperties{
		// 				Messages: []*armsupport.MessageProperties{
		// 					{
		// 						Body: to.Ptr("Hi again"),
		// 						CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 						ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
		// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-25T20:18:19.000Z"); return t}()),
		// 						Sender: to.Ptr("support engineer 2"),
		// 					},
		// 					{
		// 						Body: to.Ptr("hello"),
		// 						CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionInbound),
		// 						ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
		// 						CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-25T20:19:16.000Z"); return t}()),
		// 						Sender: to.Ptr("user"),
		// 				}},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-22T22:46:35.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
