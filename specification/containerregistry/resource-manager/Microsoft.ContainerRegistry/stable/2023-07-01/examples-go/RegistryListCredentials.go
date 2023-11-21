package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bbe1ea8bf5aa6cfbfa8855e03dbb9a93f8266bcd/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/RegistryListCredentials.json
func ExampleRegistriesClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegistriesClient().ListCredentials(ctx, "myResourceGroup", "myRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegistryListCredentialsResult = armcontainerregistry.RegistryListCredentialsResult{
	// 	Passwords: []*armcontainerregistry.RegistryPassword{
	// 		{
	// 			Name: to.Ptr(armcontainerregistry.PasswordNamePassword),
	// 			Value: to.Ptr("00000000000000000000000000000000"),
	// 		},
	// 		{
	// 			Name: to.Ptr(armcontainerregistry.PasswordNamePassword2),
	// 			Value: to.Ptr("00000000000000000000000000000000"),
	// 	}},
	// 	Username: to.Ptr("myRegistry"),
	// }
}
