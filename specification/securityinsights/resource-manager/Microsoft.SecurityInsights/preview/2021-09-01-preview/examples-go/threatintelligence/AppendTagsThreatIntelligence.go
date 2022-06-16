package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/AppendTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_AppendTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	_, err = client.AppendTags(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsight.ThreatIntelligenceAppendTags{
			ThreatIntelligenceTags: []*string{
				to.StringPtr("tag1"),
				to.StringPtr("tag2")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
