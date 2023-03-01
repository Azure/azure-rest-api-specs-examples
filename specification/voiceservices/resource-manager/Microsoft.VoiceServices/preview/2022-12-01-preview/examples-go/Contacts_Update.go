package armvoiceservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a7b696c2c73218fbca91c7c3bb625ee0f0bbea0/specification/voiceservices/resource-manager/Microsoft.VoiceServices/preview/2022-12-01-preview/examples/Contacts_Update.json
func ExampleContactsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvoiceservices.NewContactsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "testrg", "myname", "name2", armvoiceservices.ContactUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Contact = armvoiceservices.Contact{
	// 	Name: to.Ptr("name2"),
	// 	Type: to.Ptr("Microsoft.Voiceservice/communicationsGateways/contacts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.VoiceService/communicationsGateway/myname/Contacts/name2"),
	// 	Location: to.Ptr("useast"),
	// 	Properties: &armvoiceservices.ContactProperties{
	// 		ContactName: to.Ptr("John Smith"),
	// 		Email: to.Ptr("johnsmith@example.com"),
	// 		PhoneNumber: to.Ptr("+1-555-1234"),
	// 		Role: to.Ptr("Network Manager"),
	// 	},
	// }
}
