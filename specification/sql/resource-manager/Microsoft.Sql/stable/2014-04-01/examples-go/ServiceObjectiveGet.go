package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServiceObjectiveGet.json
func ExampleServiceObjectivesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceObjectivesClient().Get(ctx, "group1", "sqlcrudtest", "29dd7459-4a7c-4e56-be22-f0adda49440d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceObjective = armsql.ServiceObjective{
	// 	Name: to.Ptr("29dd7459-4a7c-4e56-be22-f0adda49440d"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/serviceObjectives"),
	// 	ID: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/serviceObjectives/29dd7459-4a7c-4e56-be22-f0adda49440d"),
	// 	Properties: &armsql.ServiceObjectiveProperties{
	// 		Enabled: to.Ptr(false),
	// 		IsDefault: to.Ptr(false),
	// 		IsSystem: to.Ptr(true),
	// 		ServiceObjectiveName: to.Ptr("System0"),
	// 	},
	// }
}
