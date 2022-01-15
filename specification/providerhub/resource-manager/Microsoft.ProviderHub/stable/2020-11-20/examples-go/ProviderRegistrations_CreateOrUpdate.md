Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.2.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armproviderhub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewProviderRegistrationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		armproviderhub.ProviderRegistration{
			Properties: &armproviderhub.ProviderRegistrationProperties{
				Capabilities: []*armproviderhub.ResourceProviderCapabilities{
					{
						Effect:  armproviderhub.ResourceProviderCapabilitiesEffect("Allow").ToPtr(),
						QuotaID: to.StringPtr("<quota-id>"),
					},
					{
						Effect:  armproviderhub.ResourceProviderCapabilitiesEffect("Allow").ToPtr(),
						QuotaID: to.StringPtr("<quota-id>"),
					}},
				Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
					IncidentContactEmail:   to.StringPtr("<incident-contact-email>"),
					IncidentRoutingService: to.StringPtr("<incident-routing-service>"),
					IncidentRoutingTeam:    to.StringPtr("<incident-routing-team>"),
				},
				ProviderType:    armproviderhub.ResourceProviderType("Internal").ToPtr(),
				ProviderVersion: to.StringPtr("<provider-version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProviderRegistrationsClientCreateOrUpdateResult)
}
```
