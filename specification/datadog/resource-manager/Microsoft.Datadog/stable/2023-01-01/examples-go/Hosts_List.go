package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Hosts_List.json
func ExampleMonitorsClient_NewListHostsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListHostsPager("myResourceGroup", "myMonitor", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.HostListResponse = armdatadog.HostListResponse{
		// 	Value: []*armdatadog.Host{
		// 		{
		// 			Name: to.Ptr("vm1"),
		// 			Aliases: []*string{
		// 				to.Ptr("vm1"),
		// 				to.Ptr("65f2dd83-95ae-4f56-b6aa-a5dafc05f4cd")},
		// 				Apps: []*string{
		// 					to.Ptr("ntp"),
		// 					to.Ptr("agent")},
		// 					Meta: &armdatadog.HostMetadata{
		// 						AgentVersion: to.Ptr("7.19.2"),
		// 						InstallMethod: &armdatadog.InstallMethod{
		// 							InstallerVersion: to.Ptr("install_script-1.0.0"),
		// 							Tool: to.Ptr("install_script"),
		// 							ToolVersion: to.Ptr("install_script"),
		// 						},
		// 						LogsAgent: &armdatadog.LogsAgent{
		// 							Transport: to.Ptr(""),
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("vm2"),
		// 					Aliases: []*string{
		// 						to.Ptr("vm2"),
		// 						to.Ptr("df631d9a-8178-4580-bf60-c697a5e8df4d")},
		// 						Apps: []*string{
		// 							to.Ptr("infra"),
		// 							to.Ptr("agent")},
		// 							Meta: &armdatadog.HostMetadata{
		// 								AgentVersion: to.Ptr("7.18.1"),
		// 								InstallMethod: &armdatadog.InstallMethod{
		// 									InstallerVersion: to.Ptr("install_script-1.0.0"),
		// 									Tool: to.Ptr("install_script"),
		// 									ToolVersion: to.Ptr("install_script"),
		// 								},
		// 								LogsAgent: &armdatadog.LogsAgent{
		// 									Transport: to.Ptr("HTTP"),
		// 								},
		// 							},
		// 					}},
		// 				}
	}
}
