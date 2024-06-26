package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetchatTranscriptDetailsForSupportTicket.json
func ExampleChatTranscriptsNoSubscriptionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChatTranscriptsNoSubscriptionClient().Get(ctx, "testticket", "b371192a-b094-4a71-b093-7246029b0a54", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChatTranscriptDetails = armsupport.ChatTranscriptDetails{
	// 	Name: to.Ptr("testmessage"),
	// 	Type: to.Ptr("Microsoft.Support/chatTranscripts"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/chatTranscripts/b371192a-b094-4a71-b093-7246029b0a54"),
	// 	Properties: &armsupport.ChatTranscriptDetailsProperties{
	// 		Messages: []*armsupport.MessageProperties{
	// 			{
	// 				Body: to.Ptr("Hi again"),
	// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 				ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
	// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-25T20:18:19.000Z"); return t}()),
	// 				Sender: to.Ptr("support engineer 2"),
	// 			},
	// 			{
	// 				Body: to.Ptr("hello"),
	// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionInbound),
	// 				ContentType: to.Ptr(armsupport.TranscriptContentType("text")),
	// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-25T20:19:16.000Z"); return t}()),
	// 				Sender: to.Ptr("user"),
	// 		}},
	// 	},
	// }
}
