package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v3"
)

// Generated from example definition: 2025-09-01/smtpUsername/getAll.json
func ExampleSMTPUsernamesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("11112222-3333-4444-5555-666677778888", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSMTPUsernamesClient().NewListPager("contosoResourceGroup", "contosoACSService", nil)
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
		// page = armcommunication.SMTPUsernamesClientListResponse{
		// 	SMTPUsernameResourceCollection: armcommunication.SMTPUsernameResourceCollection{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/communicationServices/contosoACSService/smtpUsernames?skiptoken=dae9w8wieyo29sadiue2=="),
		// 		Value: []*armcommunication.SMTPUsernameResource{
		// 			{
		// 				Name: to.Ptr("smtpusername1"),
		// 				Type: to.Ptr("Microsoft.Communication/CommunicationServices/SmtpUsername"),
		// 				ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/CommunicationServices/contosoACSService/SmtpUsername/smtpusername1"),
		// 				Properties: &armcommunication.SMTPUsernameProperties{
		// 					EntraApplicationID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa1111abcd"),
		// 					TenantID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 					Username: to.Ptr("newuser1@contoso.com"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("smtpusername2"),
		// 				Type: to.Ptr("Microsoft.Communication/CommunicationServices/SmtpUsername"),
		// 				ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/CommunicationServices/contosoACSService/SmtpUsername/smtpusername2"),
		// 				Properties: &armcommunication.SMTPUsernameProperties{
		// 					EntraApplicationID: to.Ptr("bbbb1111-1234-2222-3333-aaaa1111abcd"),
		// 					TenantID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 					Username: to.Ptr("newuser2@contoso.com"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
