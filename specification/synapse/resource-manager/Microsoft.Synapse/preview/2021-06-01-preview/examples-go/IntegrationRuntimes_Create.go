package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Create.json
func ExampleIntegrationRuntimesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginCreate(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", armsynapse.IntegrationRuntimeResource{
		Properties: &armsynapse.SelfHostedIntegrationRuntime{
			Type:        to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
			Description: to.Ptr("A selfhosted integration runtime"),
		},
	}, &armsynapse.IntegrationRuntimesClientBeginCreateOptions{IfMatch: nil})
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
	// res.IntegrationRuntimeResource = armsynapse.IntegrationRuntimeResource{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/integrationruntimes"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Synapse/workspaces/exampleWorkspaceName/integrationruntimes/exampleIntegrationRuntime"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	Properties: &armsynapse.SelfHostedIntegrationRuntime{
	// 		Type: to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
	// 		Description: to.Ptr("A selfhosted integration runtime"),
	// 	},
	// }
}
