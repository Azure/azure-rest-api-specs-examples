package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-11-02-preview/examples/ListGuardrailsVersions.json
func ExampleManagedClustersClient_NewListGuardrailsVersionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListGuardrailsVersionsPager("location1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.GuardrailsAvailableVersionsList = armcontainerservice.GuardrailsAvailableVersionsList{
		// 	Value: []*armcontainerservice.GuardrailsAvailableVersion{
		// 		{
		// 			Name: to.Ptr("v1.0.0"),
		// 			Type: to.Ptr("Microsoft.ContainerService/locations/guardrailsVersions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ContainerService/locations/location1/guardrailsVersions/v1.0.0"),
		// 			Properties: &armcontainerservice.GuardrailsAvailableVersionsProperties{
		// 				IsDefaultVersion: to.Ptr(true),
		// 				Support: to.Ptr(armcontainerservice.GuardrailsSupportPreview),
		// 			},
		// 	}},
		// }
	}
}
