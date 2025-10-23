package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2025-08-01/Grafana_FetchAvailablePlugins.json
func ExampleGrafanaClient_FetchAvailablePlugins() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGrafanaClient().FetchAvailablePlugins(ctx, "myResourceGroup", "myWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.GrafanaClientFetchAvailablePluginsResponse{
	// 	GrafanaAvailablePluginListResponse: &armdashboard.GrafanaAvailablePluginListResponse{
	// 		Value: []*armdashboard.GrafanaAvailablePlugin{
	// 			{
	// 				Name: to.Ptr("Plugin A"),
	// 				PluginID: to.Ptr("plugin-a"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Plugin B"),
	// 				PluginID: to.Ptr("plugin-b"),
	// 			},
	// 		},
	// 	},
	// }
}
