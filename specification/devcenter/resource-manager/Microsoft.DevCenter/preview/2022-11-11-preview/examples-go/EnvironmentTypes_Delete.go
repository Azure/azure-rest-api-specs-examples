package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c583b05741fadfdca116be3b9ccb1c4be8a73258/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/EnvironmentTypes_Delete.json
func ExampleEnvironmentTypesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEnvironmentTypesClient().Delete(ctx, "rg1", "Contoso", "DevTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
