package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostSecurityConnectorGovernanceRule_example.json
func ExampleGovernanceRulesClient_BeginRuleIDExecuteSingleSecurityConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewGovernanceRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRuleIDExecuteSingleSecurityConnector(ctx, "gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", &armsecurity.GovernanceRulesClientBeginRuleIDExecuteSingleSecurityConnectorOptions{ExecuteGovernanceRuleParams: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
