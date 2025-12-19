package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/ScriptExecutions_Get.json
func ExampleScriptExecutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptExecutionsClient().Get(ctx, "group1", "cloud1", "addSsoServer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.ScriptExecutionsClientGetResponse{
	// 	ScriptExecution: &armavs.ScriptExecution{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptExecutions/addSsoServer"),
	// 		Name: to.Ptr("addSsoServer"),
	// 		Properties: &armavs.ScriptExecutionProperties{
	// 			ScriptCmdletID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
	// 			Timeout: to.Ptr("P0Y0M0DT0H60M60S"),
	// 			Retention: to.Ptr("P0Y0M60DT0H60M60S"),
	// 			SubmittedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-21T17:31:28Z"); return t}()),
	// 			StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-21T17:32:28Z"); return t}()),
	// 			FinishedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-21T18:32:28Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armavs.ScriptExecutionProvisioningStateSucceeded),
	// 			FailureReason: to.Ptr("vCenter failed to connect to the external server"),
	// 			Parameters: []armavs.ScriptExecutionParameterClassification{
	// 				&armavs.ScriptStringExecutionParameter{
	// 					Name: to.Ptr("DomainName"),
	// 					Type: to.Ptr(armavs.ScriptExecutionParameterTypeValue),
	// 					Value: to.Ptr("placeholderDomain.local"),
	// 				},
	// 				&armavs.ScriptStringExecutionParameter{
	// 					Name: to.Ptr("BaseUserDN"),
	// 					Type: to.Ptr(armavs.ScriptExecutionParameterTypeValue),
	// 					Value: to.Ptr("DC=placeholder, DC=placeholder"),
	// 				},
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/scriptExecutions"),
	// 	},
	// }
}
