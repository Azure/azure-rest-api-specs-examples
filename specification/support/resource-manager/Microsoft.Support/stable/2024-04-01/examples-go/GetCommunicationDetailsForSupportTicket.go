package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetCommunicationDetailsForSupportTicket.json
func ExampleCommunicationsNoSubscriptionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunicationsNoSubscriptionClient().Get(ctx, "testticket", "testmessage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunicationDetails = armsupport.CommunicationDetails{
	// 	Name: to.Ptr("testmessage"),
	// 	Type: to.Ptr("Microsoft.Support/communications"),
	// 	ID: to.Ptr("/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage"),
	// 	Properties: &armsupport.CommunicationDetailsProperties{
	// 		Body: to.Ptr("this is a test message"),
	// 		CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 		CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-24T20:18:19.000Z"); return t}()),
	// 		Sender: to.Ptr("user@contoso.com"),
	// 		Subject: to.Ptr("this is a test message"),
	// 	},
	// }
}
