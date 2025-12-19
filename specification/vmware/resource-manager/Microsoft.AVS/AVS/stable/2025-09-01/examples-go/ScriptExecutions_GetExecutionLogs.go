package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/ScriptExecutions_GetExecutionLogs.json
func ExampleScriptExecutionsClient_GetExecutionLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptExecutionsClient().GetExecutionLogs(ctx, "group1", "cloud1", "addSsoServer", &armavs.ScriptExecutionsClientGetExecutionLogsOptions{
		ScriptOutputStreamType: []*armavs.ScriptOutputStreamType{
			to.Ptr(armavs.ScriptOutputStreamTypeInformation),
			to.Ptr(armavs.ScriptOutputStreamType("Warnings")),
			to.Ptr(armavs.ScriptOutputStreamType("Errors")),
			to.Ptr(armavs.ScriptOutputStreamTypeOutput),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.ScriptExecutionsClientGetExecutionLogsResponse{
	// 	ScriptExecution: &armavs.ScriptExecution{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptExecutions/addSsoServer"),
	// 		Name: to.Ptr("addSsoServer"),
	// 		Properties: &armavs.ScriptExecutionProperties{
	// 			Timeout: to.Ptr("P0Y0M0D0H060M0S"),
	// 			Output: []*string{
	// 				to.Ptr("Most recent output"),
	// 				to.Ptr("Second most recent output"),
	// 			},
	// 			Errors: []*string{
	// 				to.Ptr("Most recent error output"),
	// 				to.Ptr("Second most error recent output"),
	// 			},
	// 			Warnings: []*string{
	// 				to.Ptr("Most recent warning output"),
	// 				to.Ptr("Second most recent warning output"),
	// 			},
	// 			Information: []*string{
	// 				to.Ptr("Most recent information output"),
	// 				to.Ptr("Second most recent information output"),
	// 			},
	// 		},
	// 	},
	// }
}
