package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9a65970ef1837c0af1800c906aa365ba95871b26/specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_CreateOrUpdate.json
func ExampleManagedEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testcontainerenv", armappcontainers.ManagedEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.ManagedEnvironmentProperties{
			AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
				LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
					CustomerID: to.Ptr("string"),
					SharedKey:  to.Ptr("string"),
				},
			},
			DaprAIConnectionString: to.Ptr("InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/"),
			ZoneRedundant:          to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedEnvironment = armappcontainers.ManagedEnvironment{
	// 	Name: to.Ptr("testcontainerenv"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.ManagedEnvironmentProperties{
	// 		AppLogsConfiguration: &armappcontainers.AppLogsConfiguration{
	// 			LogAnalyticsConfiguration: &armappcontainers.LogAnalyticsConfiguration{
	// 				CustomerID: to.Ptr("string"),
	// 			},
	// 		},
	// 		DefaultDomain: to.Ptr("testcontainerenv.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.EnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 		ZoneRedundant: to.Ptr(true),
	// 	},
	// }
}
