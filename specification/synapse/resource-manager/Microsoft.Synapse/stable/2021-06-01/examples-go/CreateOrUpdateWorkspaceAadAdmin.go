package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspaceAadAdmin.json
func ExampleWorkspaceSQLAADAdminsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspaceSQLAADAdminsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroup1",
		"workspace1",
		armsynapse.WorkspaceAADAdminInfo{
			Properties: &armsynapse.AADAdminProperties{
				AdministratorType: to.Ptr("ActiveDirectory"),
				Login:             to.Ptr("bob@contoso.com"),
				Sid:               to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
				TenantID:          to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
			},
		},
		nil)
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
