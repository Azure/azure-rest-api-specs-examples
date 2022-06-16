package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/DeleteTenantConfiguration.json
func ExampleTenantConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armportal.NewTenantConfigurationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		armportal.ConfigurationNameDefault,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
