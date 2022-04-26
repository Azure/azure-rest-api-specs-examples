Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.4.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armproviderhub.NewProviderRegistrationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		armproviderhub.ProviderRegistration{
			Properties: &armproviderhub.ProviderRegistrationProperties{
				Capabilities: []*armproviderhub.ResourceProviderCapabilities{
					{
						Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
						QuotaID: to.Ptr("<quota-id>"),
					},
					{
						Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
						QuotaID: to.Ptr("<quota-id>"),
					}},
				Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
					IncidentContactEmail:   to.Ptr("<incident-contact-email>"),
					IncidentRoutingService: to.Ptr("<incident-routing-service>"),
					IncidentRoutingTeam:    to.Ptr("<incident-routing-team>"),
				},
				ProviderType:    to.Ptr(armproviderhub.ResourceProviderTypeInternal),
				ProviderVersion: to.Ptr("<provider-version>"),
			},
		},
		&armproviderhub.ProviderRegistrationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
