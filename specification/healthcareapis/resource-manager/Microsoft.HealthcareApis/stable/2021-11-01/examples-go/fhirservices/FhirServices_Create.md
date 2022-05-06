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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Create.json
func ExampleFhirServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewFhirServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<fhir-service-name>",
		armhealthcareapis.FhirService{
			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
				Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
			},
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
			Kind: to.Ptr(armhealthcareapis.FhirServiceKindFhirR4),
			Properties: &armhealthcareapis.FhirServiceProperties{
				AccessPolicies: []*armhealthcareapis.FhirServiceAccessPolicyEntry{
					{
						ObjectID: to.Ptr("<object-id>"),
					},
					{
						ObjectID: to.Ptr("<object-id>"),
					}},
				AcrConfiguration: &armhealthcareapis.FhirServiceAcrConfiguration{
					LoginServers: []*string{
						to.Ptr("test1.azurecr.io")},
				},
				AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
					Audience:          to.Ptr("<audience>"),
					Authority:         to.Ptr("<authority>"),
					SmartProxyEnabled: to.Ptr(true),
				},
				CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
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
				ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
					StorageAccountName: to.Ptr("<storage-account-name>"),
				},
			},
		},
		&armhealthcareapis.FhirServicesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
