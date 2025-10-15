package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2025-01-15-preview/examples/PromoteLinuxHadoopScriptAction.json
func ExampleScriptExecutionHistoryClient_Promote() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScriptExecutionHistoryClient().Promote(ctx, "rg1", "cluster1", "391145124054712", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
