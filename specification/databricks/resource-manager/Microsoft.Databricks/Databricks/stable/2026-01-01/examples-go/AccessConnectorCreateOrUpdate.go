package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks/v2"
)

// Generated from example definition: 2026-01-01/AccessConnectorCreateOrUpdate.json
func ExampleAccessConnectorsClient_BeginCreateOrUpdate_createAnAzureDatabricksAccessConnectorWithSystemAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessConnectorsClient().BeginCreateOrUpdate(ctx, "rg", "myAccessConnector", armdatabricks.AccessConnector{
		Location: to.Ptr("westus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabricks.AccessConnectorsClientCreateOrUpdateResponse{
	// 	AccessConnector: armdatabricks.AccessConnector{
	// 		Name: to.Ptr("myAccessConnector"),
	// 		Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
	// 		ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg/providers/Microsoft.Databricks/accessConnectors/myAccessConnector2"),
	// 		Identity: &armdatabricks.ManagedServiceIdentity{
	// 			Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("5619ff16-afe1-47e5-ae67-8393c6c3223d"),
	// 			TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armdatabricks.AccessConnectorProperties{
	// 			ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
	// 			ReferedBy: []*string{
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
