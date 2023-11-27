package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetCommunicationDetailsForSubscriptionSupportTicket.json
func ExampleCommunicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunicationsClient().Get(ctx, "testticket", "testmessage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunicationDetails = armsupport.CommunicationDetails{
	// 	Name: to.Ptr("testmessage"),
	// 	Type: to.Ptr("Microsoft.Support/communications"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage"),
	// 	Properties: &armsupport.CommunicationDetailsProperties{
	// 		Body: to.Ptr("this is a test message"),
	// 		CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 		CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19.000Z"); return t}()),
	// 		Sender: to.Ptr("user@contoso.com"),
	// 		Subject: to.Ptr("this is a test message"),
	// 	},
	// }
}
