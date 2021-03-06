package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PromoteLinuxHadoopScriptAction.json
func ExampleScriptExecutionHistoryClient_Promote() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewScriptExecutionHistoryClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Promote(ctx,
		"rg1",
		"cluster1",
		"391145124054712",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
