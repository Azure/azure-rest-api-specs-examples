package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/GetConfigurations.json
func ExampleLinkerClient_ListConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkerClient().ListConfigurations(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.App/containerApps/test-app", "linkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationResult = armservicelinker.ConfigurationResult{
	// 	Configurations: []*armservicelinker.SourceConfiguration{
	// 		{
	// 			Name: to.Ptr("AZURE_POSTGRESQL_HOST"),
	// 			ConfigType: to.Ptr(armservicelinker.LinkerConfigurationTypeDefault),
	// 			Value: to.Ptr("Host"),
	// 		},
	// 		{
	// 			Name: to.Ptr("AZURE_POSTGRESQL_USER"),
	// 			ConfigType: to.Ptr(armservicelinker.LinkerConfigurationTypeDefault),
	// 			Value: to.Ptr("Username"),
	// 		},
	// 		{
	// 			Name: to.Ptr("AZURE_POSTGRESQL_DATABASE"),
	// 			ConfigType: to.Ptr(armservicelinker.LinkerConfigurationTypeDefault),
	// 			Value: to.Ptr("DatabaseName"),
	// 		},
	// 		{
	// 			Name: to.Ptr("AZURE_POSTGRESQL_PORT"),
	// 			ConfigType: to.Ptr(armservicelinker.LinkerConfigurationTypeDefault),
	// 			Value: to.Ptr("Port"),
	// 		},
	// 		{
	// 			Name: to.Ptr("AZURE_POSTGRESQL_PASSWORD"),
	// 			ConfigType: to.Ptr(armservicelinker.LinkerConfigurationTypeKeyVaultSecret),
	// 			KeyVaultReferenceIdentity: to.Ptr("system"),
	// 			Value: to.Ptr("SecretUri"),
	// 	}},
	// }
}
