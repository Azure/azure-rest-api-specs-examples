package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/regenerateAgentRegistrationKey.json
func ExampleAgentRegistrationInformationClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentRegistrationInformationClient().RegenerateKey(ctx, "rg", "myAutomationAccount18", armautomation.AgentRegistrationRegenerateKeyParameter{
		KeyName: to.Ptr(armautomation.AgentRegistrationKeyNamePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.AgentRegistrationInformationClientRegenerateKeyResponse{
	// 	AgentRegistration: armautomation.AgentRegistration{
	// 		Endpoint: to.Ptr("https://eus2-agentservice-prod-1.azure-automation.net/accounts/bd8fac9e-0000-0000-0000-0000f474fbf6"),
	// 		Keys: &armautomation.AgentRegistrationKeys{
	// 			Primary: to.Ptr("5ci0000000000000000000000000000000000000000000000000000000000000000000000000000Y5H/8wFg=="),
	// 			Secondary: to.Ptr("rVp0000000000000000000000000000000000000000000000000000000000000000000000000000f8cbmrOA=="),
	// 		},
	// 	},
	// }
}
