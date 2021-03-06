package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_GetExecutionLogs.json
func ExampleScriptExecutionsClient_GetExecutionLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewScriptExecutionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetExecutionLogs(ctx,
		"group1",
		"cloud1",
		"addSsoServer",
		&armavs.ScriptExecutionsClientGetExecutionLogsOptions{ScriptOutputStreamType: []*armavs.ScriptOutputStreamType{
			to.Ptr(armavs.ScriptOutputStreamTypeInformation),
			to.Ptr(armavs.ScriptOutputStreamType("Warnings")),
			to.Ptr(armavs.ScriptOutputStreamType("Errors")),
			to.Ptr(armavs.ScriptOutputStreamTypeOutput)},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
