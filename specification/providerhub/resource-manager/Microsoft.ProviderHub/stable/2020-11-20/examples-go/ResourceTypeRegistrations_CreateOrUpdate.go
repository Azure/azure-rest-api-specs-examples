package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_CreateOrUpdate.json
func ExampleResourceTypeRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armproviderhub.NewResourceTypeRegistrationsClient("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"Microsoft.Contoso",
		"employees",
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
						SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
					}},
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
