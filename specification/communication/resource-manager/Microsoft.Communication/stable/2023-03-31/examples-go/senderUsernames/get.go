package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/senderUsernames/get.json
func ExampleSenderUsernamesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSenderUsernamesClient().Get(ctx, "contosoResourceGroup", "contosoEmailService", "contoso.com", "contosoNewsAlerts", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SenderUsernameResource = armcommunication.SenderUsernameResource{
	// 	Name: to.Ptr("contoso.com"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices/Domains/SenderUsernames"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/EmailServices/contosoEmailService/Domains/contoso.com/senderUsernames"),
	// 	Properties: &armcommunication.SenderUsernameProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		DisplayName: to.Ptr("Contoso News Alerts"),
	// 		Username: to.Ptr("contosoNewsAlerts"),
	// 	},
	// }
}
