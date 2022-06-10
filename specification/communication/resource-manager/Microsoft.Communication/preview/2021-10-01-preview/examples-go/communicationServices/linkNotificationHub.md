```go
package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/linkNotificationHub.json
func ExampleServicesClient_LinkNotificationHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcommunication.NewServicesClient("12345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.LinkNotificationHub(ctx,
		"MyResourceGroup",
		"MyCommunicationResource",
		&armcommunication.ServicesClientLinkNotificationHubOptions{LinkNotificationHubParameters: &armcommunication.LinkNotificationHubParameters{
			ConnectionString: to.Ptr("Endpoint=sb://MyNamespace.servicebus.windows.net/;SharedAccessKey=abcd1234"),
			ResourceID:       to.Ptr("/subscriptions/12345/resourceGroups/MyOtherResourceGroup/providers/Microsoft.NotificationHubs/namespaces/MyNamespace/notificationHubs/MyHub"),
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcommunication%2Farmcommunication%2Fv2.0.0-beta.1/sdk/resourcemanager/communication/armcommunication/README.md) on how to add the SDK to your project and authenticate.
