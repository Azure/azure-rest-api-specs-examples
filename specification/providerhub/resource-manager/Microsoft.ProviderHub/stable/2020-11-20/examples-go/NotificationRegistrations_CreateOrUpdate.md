Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.4.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_CreateOrUpdate.json
func ExampleNotificationRegistrationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armproviderhub.NewNotificationRegistrationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<provider-namespace>",
		"<notification-registration-name>",
		armproviderhub.NotificationRegistration{
			Properties: &armproviderhub.NotificationRegistrationProperties{
				IncludedEvents: []*string{
					to.Ptr("*/write"),
					to.Ptr("Microsoft.Contoso/employees/delete")},
				MessageScope: to.Ptr(armproviderhub.MessageScopeRegisteredSubscriptions),
				NotificationEndpoints: []*armproviderhub.NotificationEndpoint{
					{
						Locations: []*string{
							to.Ptr(""),
							to.Ptr("East US")},
						NotificationDestination: to.Ptr("<notification-destination>"),
					},
					{
						Locations: []*string{
							to.Ptr("North Europe")},
						NotificationDestination: to.Ptr("<notification-destination>"),
					}},
				NotificationMode: to.Ptr(armproviderhub.NotificationModeEventHub),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
