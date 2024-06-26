package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
func ExampleWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient().Get(ctx, "workspace-6852", "workspace-2080", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedSQLminimalTLSSettings = armsynapse.DedicatedSQLminimalTLSSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/dedicatedSQLminimalTlsSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-6852/providers/Microsoft.Synapse/workspaces/workspace-2080/dedicatedSQLminimalTlsSettings/default"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armsynapse.DedicatedSQLminimalTLSSettingsProperties{
	// 		MinimalTLSVersion: to.Ptr("1.1"),
	// 	},
	// }
}
