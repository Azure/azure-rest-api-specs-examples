package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/89260be1a92c914b7b48af8e8f75938d5e76851d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-08-02-preview/examples/GetGuardrailsVersions.json
func ExampleManagedClustersClient_GetGuardrailsVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetGuardrailsVersions(ctx, "location1", "v1.0.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GuardrailsAvailableVersion = armcontainerservice.GuardrailsAvailableVersion{
	// 	Name: to.Ptr("v1.0.0"),
	// 	Type: to.Ptr("Microsoft.ContainerService/locations/guardrailsVersions"),
	// 	ID: to.Ptr("/subscriptions/subid1/providers/Microsoft.ContainerService/locations/location1/guardrailsVersions/v1.0.0"),
	// 	Properties: &armcontainerservice.GuardrailsAvailableVersionsProperties{
	// 		IsDefaultVersion: to.Ptr(true),
	// 		Support: to.Ptr(armcontainerservice.GuardrailsSupportPreview),
	// 	},
	// }
}
