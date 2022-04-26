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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_CreateOrUpdate.json
func ExampleResourceTypeRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		"<resource-type>",
		armproviderhub.ResourceTypeRegistration{
			Properties: &armproviderhub.ResourceTypeRegistrationProperties{
				Endpoints: []*armproviderhub.ResourceTypeEndpoint{
					{
						APIVersions: []*string{
							to.Ptr("2020-06-01-preview")},
						Locations: []*string{
							to.Ptr("West US"),
							to.Ptr("East US"),
							to.Ptr("North Europe")},
						RequiredFeatures: []*string{
							to.Ptr("<feature flag>")},
					}},
				Regionality: to.Ptr(armproviderhub.RegionalityRegional),
				RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
				SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
					{
						APIVersions: []*string{
							to.Ptr("2020-06-01-preview")},
						SwaggerSpecFolderURI: to.Ptr("<swagger-spec-folder-uri>"),
					}},
			},
		},
		&armproviderhub.ResourceTypeRegistrationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
