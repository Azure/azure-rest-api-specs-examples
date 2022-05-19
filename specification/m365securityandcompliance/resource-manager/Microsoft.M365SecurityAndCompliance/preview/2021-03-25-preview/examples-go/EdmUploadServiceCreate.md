Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fm365securityandcompliance%2Farmm365securityandcompliance%2Fv0.5.0/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance/README.md) on how to add the SDK to your project and authenticate.

```go
package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/EdmUploadServiceCreate.json
func ExamplePrivateLinkServicesForEDMUploadClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForEDMUploadClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"service1",
		armm365securityandcompliance.PrivateLinkServicesForEDMUploadDescription{
			Identity: &armm365securityandcompliance.ServicesResourceIdentity{
				Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
			Location: to.Ptr("westus2"),
			Tags:     map[string]*string{},
			Properties: &armm365securityandcompliance.ServicesProperties{
				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
					{
						ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
					{
						ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
					}},
				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
					Audience:          to.Ptr("https://azurehealthcareapis.com"),
					Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
					SmartProxyEnabled: to.Ptr(true),
				},
				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
					AllowCredentials: to.Ptr(false),
					Headers: []*string{
						to.Ptr("*")},
					MaxAge: to.Ptr[int64](1440),
					Methods: []*string{
						to.Ptr("DELETE"),
						to.Ptr("GET"),
						to.Ptr("OPTIONS"),
						to.Ptr("PATCH"),
						to.Ptr("POST"),
						to.Ptr("PUT")},
					Origins: []*string{
						to.Ptr("*")},
				},
				CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
					OfferThroughput: to.Ptr[int64](1000),
				},
				ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
					StorageAccountName: to.Ptr("existingStorageAccount"),
				},
				PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
				PublicNetworkAccess:        to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
