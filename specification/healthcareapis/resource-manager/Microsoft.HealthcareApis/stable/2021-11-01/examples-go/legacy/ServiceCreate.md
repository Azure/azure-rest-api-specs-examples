Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.4.0/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armhealthcareapis.ServicesDescription{
			Identity: &armhealthcareapis.ServicesResourceIdentity{
				Type: to.Ptr(armhealthcareapis.ManagedServiceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr(armhealthcareapis.KindFhirR4),
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armhealthcareapis.ServicesProperties{
				AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
					{
						ObjectID: to.Ptr("<object-id>"),
					},
					{
						ObjectID: to.Ptr("<object-id>"),
					}},
				AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
					Audience:          to.Ptr("<audience>"),
					Authority:         to.Ptr("<authority>"),
					SmartProxyEnabled: to.Ptr(true),
				},
				CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
					AllowCredentials: to.Ptr(false),
					Headers: []*string{
						to.Ptr("*")},
					MaxAge: to.Ptr[int32](1440),
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
				CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.Ptr("<key-vault-key-uri>"),
					OfferThroughput: to.Ptr[int32](1000),
				},
				ExportConfiguration: &armhealthcareapis.ServiceExportConfigurationInfo{
					StorageAccountName: to.Ptr("<storage-account-name>"),
				},
				PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{},
				PublicNetworkAccess:        to.Ptr(armhealthcareapis.PublicNetworkAccessDisabled),
			},
		},
		&armhealthcareapis.ServicesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
