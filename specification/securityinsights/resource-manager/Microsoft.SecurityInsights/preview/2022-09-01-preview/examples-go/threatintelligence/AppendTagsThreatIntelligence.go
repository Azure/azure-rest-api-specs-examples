package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/AppendTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_AppendTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewThreatIntelligenceIndicatorClient().AppendTags(ctx, "myRg", "myWorkspace", "d9cd6f0b-96b9-3984-17cd-a779d1e15a93", armsecurityinsights.ThreatIntelligenceAppendTags{
		ThreatIntelligenceTags: []*string{
			to.Ptr("tag1"),
			to.Ptr("tag2")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
