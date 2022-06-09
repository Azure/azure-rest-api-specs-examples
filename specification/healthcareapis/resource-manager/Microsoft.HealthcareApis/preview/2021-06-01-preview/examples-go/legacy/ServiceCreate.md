```go
package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/legacy/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armhealthcareapis.ServicesDescription{
			Identity: &armhealthcareapis.ServicesResourceIdentity{
				Type: armhealthcareapis.ManagedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Kind:     armhealthcareapis.KindFhirR4.ToPtr(),
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armhealthcareapis.ServicesProperties{
				AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
					},
					{
						ObjectID: to.StringPtr("<object-id>"),
					}},
				AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
					Audience:          to.StringPtr("<audience>"),
					Authority:         to.StringPtr("<authority>"),
					SmartProxyEnabled: to.BoolPtr(true),
				},
				CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
					AllowCredentials: to.BoolPtr(false),
					Headers: []*string{
						to.StringPtr("*")},
					MaxAge: to.Int32Ptr(1440),
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
				CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.StringPtr("<key-vault-key-uri>"),
					OfferThroughput: to.Int32Ptr(1000),
				},
				ExportConfiguration: &armhealthcareapis.ServiceExportConfigurationInfo{
					StorageAccountName: to.StringPtr("<storage-account-name>"),
				},
				PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{},
				PublicNetworkAccess:        armhealthcareapis.PublicNetworkAccess("Disabled").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ServicesClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.2.1/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.
