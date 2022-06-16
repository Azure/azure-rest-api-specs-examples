package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_ReplaceTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	res, err := client.ReplaceTags(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsight.ThreatIntelligenceIndicatorModelForRequestBody{
			Kind: armsecurityinsight.ThreatIntelligenceResourceKindEnum("indicator").ToPtr(),
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsight.ThreatIntelligenceIndicatorProperties{
				ThreatIntelligenceTags: []*string{
					to.StringPtr("patching tags")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ThreatIntelligenceIndicatorClientReplaceTagsResult)
}
