package armmonitorworkspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitorworkspaces"
)

// Generated from example definition: 2025-10-03/Issue_FetchBackgroundVisualization_MaximumSet_Gen.json
func ExampleIssueClient_FetchBackgroundVisualization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitorworkspaces.NewClientFactory("aceaa046-91f0-492a-96dc-45e10a9183dc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIssueClient().FetchBackgroundVisualization(ctx, "rg1", "myWorkspace", "3f29e1b2b05f8371595dc761fed8e8b3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitorworkspaces.IssueClientFetchBackgroundVisualizationResponse{
	// 	BackgroundVisualization: armmonitorworkspaces.BackgroundVisualization{
	// 		Visualization: to.Ptr("{\"type\":\"AdaptiveCard\",\"version\":\"1.3\",\"body\":[{\"type\":\"TextBlock\",\"text\":\"Issue Background Visualization\",\"size\":\"Large\",\"weight\":\"Bolder\"},{\"type\":\"TextBlock\",\"text\":\"This adaptive card provides background context for the issue including related metrics, logs, and diagnostic information.\",\"wrap\":true},{\"type\":\"FactSet\",\"facts\":[{\"title\":\"Affected Resources\",\"value\":\"5 virtual machines\"},{\"title\":\"Impact Duration\",\"value\":\"2 hours 15 minutes\"},{\"title\":\"Severity Level\",\"value\":\"High\"}]},{\"type\":\"Container\",\"items\":[{\"type\":\"TextBlock\",\"text\":\"Key Metrics\",\"weight\":\"Bolder\"},{\"type\":\"TextBlock\",\"text\":\"CPU Utilization: 95%\\nMemory Usage: 87%\\nDisk I/O: High\",\"wrap\":true}]}]}"),
	// 		Origin: &armmonitorworkspaces.Origin{
	// 			AddedBy: to.Ptr("171a811c-2a3a-4e6c-b742-f78f5f6ca51c"),
	// 			AddedByType: to.Ptr(armmonitorworkspaces.AddedByTypeManual),
	// 		},
	// 	},
	// }
}
