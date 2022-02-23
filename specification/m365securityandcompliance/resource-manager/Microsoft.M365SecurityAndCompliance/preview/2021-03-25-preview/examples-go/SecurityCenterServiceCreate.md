Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fm365securityandcompliance%2Farmm365securityandcompliance%2Fv0.2.1/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceCreate.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armm365securityandcompliance.NewPrivateLinkServicesForM365SecurityCenterClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
			Identity: &armm365securityandcompliance.ServicesResourceIdentity{
				Type: armm365securityandcompliance.ManagedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Kind:     armm365securityandcompliance.KindFhirR4.ToPtr(),
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armm365securityandcompliance.ServicesProperties{
				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
					},
					{
						ObjectID: to.StringPtr("<object-id>"),
					}},
				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
					Audience:          to.StringPtr("<audience>"),
					Authority:         to.StringPtr("<authority>"),
					SmartProxyEnabled: to.BoolPtr(true),
				},
				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
					AllowCredentials: to.BoolPtr(false),
					Headers: []*string{
						to.StringPtr("*")},
					MaxAge: to.Int64Ptr(1440),
					Methods: []*string{
						to.StringPtr("DELETE"),
						to.StringPtr("GET"),
						to.StringPtr("OPTIONS"),
						to.StringPtr("PATCH"),
						to.StringPtr("POST"),
						to.StringPtr("PUT")},
					Origins: []*string{
						to.StringPtr("*")},
				},
				CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.StringPtr("<key-vault-key-uri>"),
					OfferThroughput: to.Int64Ptr(1000),
				},
				ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
					StorageAccountName: to.StringPtr("<storage-account-name>"),
				},
				PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
				PublicNetworkAccess:        armm365securityandcompliance.PublicNetworkAccess("Disabled").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkServicesForM365SecurityCenterClientCreateOrUpdateResult)
}
```
