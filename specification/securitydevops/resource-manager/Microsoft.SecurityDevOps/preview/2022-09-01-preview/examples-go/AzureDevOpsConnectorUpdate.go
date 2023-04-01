package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorUpdate.json
func ExampleAzureDevOpsConnectorClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsConnectorClient().BeginUpdate(ctx, "westusrg", "testconnector", armsecuritydevops.AzureDevOpsConnector{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"client": to.Ptr("dev-client"),
			"env":    to.Ptr("dev"),
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
	// res.AzureDevOpsConnector = armsecuritydevops.AzureDevOpsConnector{
	// 	Name: to.Ptr("testconnector1"),
	// 	Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// 	Properties: &armsecuritydevops.AzureDevOpsConnectorProperties{
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 	},
	// }
}
