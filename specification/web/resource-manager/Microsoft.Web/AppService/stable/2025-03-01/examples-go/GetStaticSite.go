package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSite.json
func ExampleStaticSitesClient_GetStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().GetStaticSite(ctx, "rg", "testStaticSite0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 			{
	// 				BackendResourceID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ApiManagement/service/apimService0"),
	// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T20:57:24.805Z"); return t}()),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 		}},
	// 		PrivateEndpointConnections: []*armappservice.ResponseMessageEnvelopeRemotePrivateEndpointConnection{
	// 		},
	// 		RepositoryURL: to.Ptr("https://github.com/username/repo"),
	// 		StagingEnvironmentPolicy: to.Ptr(armappservice.StagingEnvironmentPolicyEnabled),
	// 	},
	// 	SKU: &armappservice.SKUDescription{
	// 		Name: to.Ptr("Basic"),
	// 		Tier: to.Ptr("Basic"),
	// 	},
	// }
}
