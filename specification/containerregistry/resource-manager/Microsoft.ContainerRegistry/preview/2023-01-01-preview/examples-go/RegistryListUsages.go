package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/RegistryListUsages.json
func ExampleRegistriesClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewRegistriesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListUsages(ctx, "myResourceGroup", "myRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegistryUsageListResult = armcontainerregistry.RegistryUsageListResult{
	// 	Value: []*armcontainerregistry.RegistryUsage{
	// 		{
	// 			Name: to.Ptr("Size"),
	// 			CurrentValue: to.Ptr[int64](12345678),
	// 			Limit: to.Ptr[int64](107374182400),
	// 			Unit: to.Ptr(armcontainerregistry.RegistryUsageUnitBytes),
	// 		},
	// 		{
	// 			Name: to.Ptr("Webhooks"),
	// 			CurrentValue: to.Ptr[int64](2),
	// 			Limit: to.Ptr[int64](10),
	// 			Unit: to.Ptr(armcontainerregistry.RegistryUsageUnitCount),
	// 	}},
	// }
}
