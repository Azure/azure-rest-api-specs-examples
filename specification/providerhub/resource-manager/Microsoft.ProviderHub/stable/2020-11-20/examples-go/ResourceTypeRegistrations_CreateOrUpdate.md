Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.2.1/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_CreateOrUpdate.json
func ExampleResourceTypeRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		"<resource-type>",
		armproviderhub.ResourceTypeRegistration{
			Properties: &armproviderhub.ResourceTypeRegistrationProperties{
				Endpoints: []*armproviderhub.ResourceTypeEndpoint{
					{
						APIVersions: []*string{
							to.StringPtr("2020-06-01-preview")},
						Locations: []*string{
							to.StringPtr("West US"),
							to.StringPtr("East US"),
							to.StringPtr("North Europe")},
						RequiredFeatures: []*string{
							to.StringPtr("<feature flag>")},
					}},
				Regionality: armproviderhub.Regionality("Regional").ToPtr(),
				RoutingType: armproviderhub.RoutingType("Default").ToPtr(),
				SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
					{
						APIVersions: []*string{
							to.StringPtr("2020-06-01-preview")},
						SwaggerSpecFolderURI: to.StringPtr("<swagger-spec-folder-uri>"),
					}},
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
	log.Printf("Response result: %#v\n", res.ResourceTypeRegistrationsClientCreateOrUpdateResult)
}
```
