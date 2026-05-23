package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: 2024-04-01/GetchatTranscriptDetailsForSubscriptionSupportTicket.json
func ExampleChatTranscriptsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("132d901f-189d-4381-9214-fe68e27e05a1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChatTranscriptsClient().Get(ctx, "testticket", "69586795-45e9-45b5-bd9e-c9bb237d3e44", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsupport.ChatTranscriptsClientGetResponse{
	// 	ChatTranscriptDetails: &armsupport.ChatTranscriptDetails{
	// 		Name: to.Ptr("testmessage"),
	// 		Type: to.Ptr("Microsoft.Support/chatTranscripts"),
	// 		ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket/chatTranscripts/69586795-45e9-45b5-bd9e-c9bb237d3e44"),
	// 		Properties: &armsupport.ChatTranscriptDetailsProperties{
	// 			Messages: []*armsupport.MessageProperties{
	// 				{
	// 					Body: to.Ptr("Hi again"),
	// 					CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 					ContentType: to.Ptr("text"),
	// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-23T20:18:19Z"); return t}()),
	// 					Sender: to.Ptr("support engineer 2"),
	// 				},
	// 				{
	// 					Body: to.Ptr("hello"),
	// 					CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionInbound),
	// 					ContentType: to.Ptr("text"),
	// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-23T20:19:16Z"); return t}()),
	// 					Sender: to.Ptr("user"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
