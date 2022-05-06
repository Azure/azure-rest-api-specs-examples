Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fm365securityandcompliance%2Farmm365securityandcompliance%2Fv0.4.0/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance/README.md) on how to add the SDK to your project and authenticate.

```go
package armm365securityandcompliance_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ManagementAPIServiceCreate.json
func ExamplePrivateLinkServicesForO365ManagementActivityAPIClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForO365ManagementActivityAPIClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armm365securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIDescription{
			Identity: &armm365securityandcompliance.ServicesResourceIdentity{
				Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armm365securityandcompliance.ServicesProperties{
				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
					{
						ObjectID: to.Ptr("<object-id>"),
					},
					{
						ObjectID: to.Ptr("<object-id>"),
					}},
				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
					Audience:          to.Ptr("<audience>"),
					Authority:         to.Ptr("<authority>"),
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
					KeyVaultKeyURI:  to.Ptr("<key-vault-key-uri>"),
					OfferThroughput: to.Ptr[int64](1000),
				},
				ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
					StorageAccountName: to.Ptr("<storage-account-name>"),
				},
				PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
				PublicNetworkAccess:        to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
			},
		},
		&armm365securityandcompliance.PrivateLinkServicesForO365ManagementActivityAPIClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
