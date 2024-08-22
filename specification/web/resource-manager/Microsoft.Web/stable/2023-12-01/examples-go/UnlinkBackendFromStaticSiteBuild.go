package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/UnlinkBackendFromStaticSiteBuild.json
func ExampleStaticSitesClient_UnlinkBackendFromBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStaticSitesClient().UnlinkBackendFromBuild(ctx, "rg", "testStaticSite0", "12", "testBackend", &armappservice.StaticSitesClientUnlinkBackendFromBuildOptions{IsCleaningAuthConfig: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
