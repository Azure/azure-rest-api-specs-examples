package armintegrationspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/azureintegrationspaces/resource-manager/Microsoft.IntegrationSpaces/preview/2023-11-14-preview/examples/InfrastructureResources_Get.json
func ExampleInfrastructureResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armintegrationspaces.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInfrastructureResourcesClient().Get(ctx, "testrg", "Space1", "InfrastructureResource1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InfrastructureResource = armintegrationspaces.InfrastructureResource{
	// 	Name: to.Ptr("InfrastructureResource1"),
	// 	Type: to.Ptr("Microsoft.IntegrationSpaces/spaces/infrastructureresources"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.IntegrationSpaces/spaces/Space1/infrastructureResources/InfrastructureResource1"),
	// 	Properties: &armintegrationspaces.InfrastructureResourceProperties{
	// 		ProvisioningState: to.Ptr(armintegrationspaces.ProvisioningStateSucceeded),
	// 		ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ApiManagement/service/APIM1"),
	// 		ResourceType: to.Ptr("Microsoft.ApiManagement/service"),
	// 	},
	// }
}
