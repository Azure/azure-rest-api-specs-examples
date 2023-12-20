package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Troubleshooter_Restart.json
func ExampleTroubleshootersClient_Restart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTroubleshootersClient().Restart(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestartTroubleshooterResponse = armselfhelp.RestartTroubleshooterResponse{
	// 	TroubleshooterResourceName: to.Ptr("SampleTroubleshooterResourceName"),
	// }
}
