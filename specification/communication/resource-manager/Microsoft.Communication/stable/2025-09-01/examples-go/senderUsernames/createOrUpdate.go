package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v3"
)

// Generated from example definition: 2025-09-01/senderUsernames/createOrUpdate.json
func ExampleSenderUsernamesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("11112222-3333-4444-5555-666677778888", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSenderUsernamesClient().CreateOrUpdate(ctx, "contosoResourceGroup", "contosoEmailService", "contoso.com", "contosoNewsAlerts", armcommunication.SenderUsernameResource{
		Properties: &armcommunication.SenderUsernameProperties{
			DisplayName: to.Ptr("Contoso News Alerts"),
			Username:    to.Ptr("contosoNewsAlerts"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommunication.SenderUsernamesClientCreateOrUpdateResponse{
	// 	SenderUsernameResource: armcommunication.SenderUsernameResource{
	// 		Name: to.Ptr("contoso.com"),
	// 		Type: to.Ptr("Microsoft.Communication/EmailServices/Domains"),
	// 		ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/EmailServices/contosoEmailService/Domains/contoso.com/senderUsernames/contosoNewsAlerts"),
	// 		Properties: &armcommunication.SenderUsernameProperties{
	// 			DataLocation: to.Ptr("United States"),
	// 			DisplayName: to.Ptr("Contoso News Alerts"),
	// 			Username: to.Ptr("contosoNewsAlerts"),
	// 		},
	// 	},
	// }
}
