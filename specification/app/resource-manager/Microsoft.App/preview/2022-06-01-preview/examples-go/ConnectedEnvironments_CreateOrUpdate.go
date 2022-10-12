package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironments_CreateOrUpdate.json
func ExampleConnectedEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewConnectedEnvironmentsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "examplerg", "testenv", armappcontainers.ConnectedEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ConnectedEnvironmentProperties{
			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
				CertificatePassword: []byte("private key password"),
				CertificateValue:    []byte("PFX-or-PEM-blob"),
				DNSSuffix:           to.Ptr("www.my-name.com"),
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			StaticIP:               to.Ptr("1.2.3.4"),
		},
	}, nil)
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
