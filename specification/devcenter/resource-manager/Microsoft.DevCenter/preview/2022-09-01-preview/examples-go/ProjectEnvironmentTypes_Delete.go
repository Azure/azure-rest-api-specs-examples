package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/ProjectEnvironmentTypes_Delete.json
func ExampleProjectEnvironmentTypesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "rg1", "ContosoProj", "{environmentTypeName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
