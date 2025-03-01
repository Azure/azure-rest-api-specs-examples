package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/Targets_Get_MaximumSet_Gen.json
func ExampleTargetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().Get(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.TargetsClientGetResponse{
	// 	Target: &armdatabasewatcher.Target{
	// 		Properties: &armdatabasewatcher.SQLDbSingleDatabaseTargetProperties{
	// 			TargetType: to.Ptr("SqlDb"),
	// 			TargetAuthenticationType: to.Ptr(armdatabasewatcher.TargetAuthenticationTypeAAD),
	// 			ConnectionServerName: to.Ptr("sqlServero1ihe2"),
	// 			SQLDbResourceID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.Sql/servers/m1/databases/m2"),
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-px9ma7/providers/Microsoft.DatabaseWatcher/watchers/databasemo3d9sgt/targets/monitoringzkndgm"),
	// 		Name: to.Ptr("monitoringzkndgm"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers/targets"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 			CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 			CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mxp"),
	// 			LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		},
	// 	},
	// }
}
