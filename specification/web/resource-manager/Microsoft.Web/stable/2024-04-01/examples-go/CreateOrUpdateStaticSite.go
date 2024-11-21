package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/CreateOrUpdateStaticSite.json
func ExampleStaticSitesClient_BeginCreateOrUpdateStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStaticSitesClient().BeginCreateOrUpdateStaticSite(ctx, "rg", "testStaticSite0", armappservice.StaticSiteARMResource{
		Location: to.Ptr("West US 2"),
		Properties: &armappservice.StaticSite{
			Branch: to.Ptr("master"),
			BuildProperties: &armappservice.StaticSiteBuildProperties{
				APILocation:         to.Ptr("api"),
				AppArtifactLocation: to.Ptr("build"),
				AppLocation:         to.Ptr("app"),
			},
			RepositoryToken: to.Ptr("repoToken123"),
			RepositoryURL:   to.Ptr("https://github.com/username/RepoName"),
		},
		SKU: &armappservice.SKUDescription{
			Name: to.Ptr("Basic"),
			Tier: to.Ptr("Basic"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StaticSiteARMResource = armappservice.StaticSiteARMResource{
	// 	Name: to.Ptr("testStaticSite0"),
	// 	Type: to.Ptr("Microsoft.Web/staticSites"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0"),
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armappservice.StaticSite{
	// 		AllowConfigFileUpdates: to.Ptr(true),
	// 		Branch: to.Ptr("demo"),
	// 		ContentDistributionEndpoint: to.Ptr(""),
	// 		CustomDomains: []*string{
	// 		},
	// 		DefaultHostname: to.Ptr("happy-sea-15afae3e.azurestaticwebsites.net"),
	// 		KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
	// 		LinkedBackends: []*armappservice.StaticSiteLinkedBackend{
	// 		},
	// 		PrivateEndpointConnections: []*armappservice.ResponseMessageEnvelopeRemotePrivateEndpointConnection{
	// 		},
	// 		RepositoryURL: to.Ptr("https://github.com/username/RepoName"),
	// 		StagingEnvironmentPolicy: to.Ptr(armappservice.StagingEnvironmentPolicyEnabled),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("Basic"),
	// 		Tier: to.Ptr("Basic"),
	// 	},
	// }
}
