package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorPatchUpdate.json
func ExampleAccessConnectorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessConnectorsClient().BeginUpdate(ctx, "rg", "myAccessConnector", armdatabricks.AccessConnectorUpdate{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
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
	// res.AccessConnector = armdatabricks.AccessConnector{
	// 	Name: to.Ptr("myAccessConnector"),
	// 	Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Identity: &armdatabricks.ManagedServiceIdentity{
	// 		Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("5619ff16-afe1-47e5-ae67-8393c6c3223d"),
	// 		TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
	// 	},
	// 	Properties: &armdatabricks.AccessConnectorProperties{
	// 		ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
	// 	},
	// }
}
