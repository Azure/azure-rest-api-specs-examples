Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv1.0.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armproviderhub.NewProviderRegistrationsClient("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"Microsoft.Contoso",
		armproviderhub.ProviderRegistration{
			Properties: &armproviderhub.ProviderRegistrationProperties{
				Capabilities: []*armproviderhub.ResourceProviderCapabilities{
					{
						Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
						QuotaID: to.Ptr("CSP_2015-05-01"),
					},
					{
						Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
						QuotaID: to.Ptr("CSP_MG_2017-12-01"),
					}},
				Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
					IncidentContactEmail:   to.Ptr("helpme@contoso.com"),
					IncidentRoutingService: to.Ptr("Contoso Resource Provider"),
					IncidentRoutingTeam:    to.Ptr("Contoso Triage"),
				},
				ProviderType:    to.Ptr(armproviderhub.ResourceProviderTypeInternal),
				ProviderVersion: to.Ptr("2.0"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
