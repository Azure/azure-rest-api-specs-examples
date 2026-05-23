package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: 2024-04-01/CreateSupportTicketCommunication.json
func ExampleCommunicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("132d901f-189d-4381-9214-fe68e27e05a1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunicationsClient().BeginCreate(ctx, "testticket", "testcommunication", armsupport.CommunicationDetails{
		Properties: &armsupport.CommunicationDetailsProperties{
			Body:    to.Ptr("This is a test message from a customer!"),
			Sender:  to.Ptr("user@contoso.com"),
			Subject: to.Ptr("This is a test message from a customer!"),
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
	// res = armsupport.CommunicationsClientCreateResponse{
	// 	CommunicationDetails: &armsupport.CommunicationDetails{
	// 		Name: to.Ptr("testcommunication"),
	// 		Type: to.Ptr("Microsoft.Support/communications"),
	// 		ID: to.Ptr("/subscriptions/132d901f-189d-4381-9214-fe68e27e05a1/providers/Microsoft.Support/supportTickets/testticket/communications/testcommunication"),
	// 		Properties: &armsupport.CommunicationDetailsProperties{
	// 			Body: to.Ptr("This is a test message from a customer!"),
	// 			CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 			CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19Z"); return t}()),
	// 			Sender: to.Ptr("user@contoso.com"),
	// 			Subject: to.Ptr("This is a test message from a customer!"),
	// 		},
	// 	},
	// }
}
