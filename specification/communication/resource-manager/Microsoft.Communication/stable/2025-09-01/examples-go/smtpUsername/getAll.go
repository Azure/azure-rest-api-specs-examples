package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7832c9e47b8998a1c994b98345eea24dbc2ac5b8/specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/getAll.json
func ExampleSMTPUsernamesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.SMTPUsernameResourceCollection = armcommunication.SMTPUsernameResourceCollection{
		// 	Value: []*armcommunication.SMTPUsernameResource{
		// 		{
		// 			Name: to.Ptr("smtpusername1"),
		// 			Type: to.Ptr("Microsoft.Communication/CommunicationServices/SmtpUsername"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/CommunicationServices/contosoACSService/SmtpUsername/smtpusername1"),
		// 			Properties: &armcommunication.SMTPUsernameProperties{
		// 				EntraApplicationID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa1111abcd"),
		// 				TenantID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 				Username: to.Ptr("newuser1@contoso.com"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("smtpusername2"),
		// 			Type: to.Ptr("Microsoft.Communication/CommunicationServices/SmtpUsername"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/CommunicationServices/contosoACSService/SmtpUsername/smtpusername2"),
		// 			Properties: &armcommunication.SMTPUsernameProperties{
		// 				EntraApplicationID: to.Ptr("bbbb1111-1234-2222-3333-aaaa1111abcd"),
		// 				TenantID: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 				Username: to.Ptr("newuser2@contoso.com"),
		// 			},
		// 	}},
		// }
	}
}
