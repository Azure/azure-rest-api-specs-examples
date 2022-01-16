Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.2.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_SendMessage_Post.json
func ExampleUserSessionsClient_SendMessage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewUserSessionsClient("<subscription-id>", cred, nil)
	_, err = client.SendMessage(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		"<session-host-name>",
		"<user-session-id>",
		&armdesktopvirtualization.UserSessionsClientSendMessageOptions{SendMessage: &armdesktopvirtualization.SendMessage{
			MessageBody:  to.StringPtr("<message-body>"),
			MessageTitle: to.StringPtr("<message-title>"),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
```
