Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.2.1/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/fhirservices/FhirServices_Create.json
func ExampleFhirServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<fhir-service-name>",
		armhealthcareapis.FhirService{
			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
				Type: armhealthcareapis.ManagedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"additionalProp1": to.StringPtr("string"),
				"additionalProp2": to.StringPtr("string"),
				"additionalProp3": to.StringPtr("string"),
			},
			Kind: armhealthcareapis.FhirServiceKind("fhir-R4").ToPtr(),
			Properties: &armhealthcareapis.FhirServiceProperties{
				AccessPolicies: []*armhealthcareapis.FhirServiceAccessPolicyEntry{
					{
						ObjectID: to.StringPtr("<object-id>"),
					},
					{
						ObjectID: to.StringPtr("<object-id>"),
					}},
				AcrConfiguration: &armhealthcareapis.FhirServiceAcrConfiguration{
					LoginServers: []*string{
						to.StringPtr("test1.azurecr.io")},
				},
				AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
					Audience:          to.StringPtr("<audience>"),
					Authority:         to.StringPtr("<authority>"),
					SmartProxyEnabled: to.BoolPtr(true),
				},
				CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
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
				ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
					StorageAccountName: to.StringPtr("<storage-account-name>"),
				},
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
	log.Printf("Response result: %#v\n", res.FhirServicesClientCreateOrUpdateResult)
}
```
