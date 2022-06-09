```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_CreateIndicator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	res, err := client.CreateIndicator(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsecurityinsight.ThreatIntelligenceIndicatorModelForRequestBody{
			Kind: armsecurityinsight.ThreatIntelligenceResourceKindEnum("indicator").ToPtr(),
			Properties: &armsecurityinsight.ThreatIntelligenceIndicatorProperties{
				Description:        to.StringPtr("<description>"),
				Confidence:         to.Int32Ptr(78),
				CreatedByRef:       to.StringPtr("<created-by-ref>"),
				DisplayName:        to.StringPtr("<display-name>"),
				ExternalReferences: []*armsecurityinsight.ThreatIntelligenceExternalReference{},
				GranularMarkings:   []*armsecurityinsight.ThreatIntelligenceGranularMarkingModel{},
				KillChainPhases:    []*armsecurityinsight.ThreatIntelligenceKillChainPhase{},
				Labels:             []*string{},
				Modified:           to.StringPtr("<modified>"),
				Pattern:            to.StringPtr("<pattern>"),
				PatternType:        to.StringPtr("<pattern-type>"),
				Revoked:            to.BoolPtr(false),
				Source:             to.StringPtr("<source>"),
				ThreatIntelligenceTags: []*string{
					to.StringPtr("new schema")},
				ThreatTypes: []*string{
					to.StringPtr("compromised")},
				ValidFrom:  to.StringPtr("<valid-from>"),
				ValidUntil: to.StringPtr("<valid-until>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ThreatIntelligenceIndicatorClientCreateIndicatorResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.
