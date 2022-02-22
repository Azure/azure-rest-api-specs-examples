Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.1.1/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_CreateIndicator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	res, err := client.CreateIndicator(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModelForRequestBody{
			Kind: armsecurityinsights.ThreatIntelligenceResourceKindEnum("indicator").ToPtr(),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				Description:        to.StringPtr("<description>"),
				Confidence:         to.Int32Ptr(78),
				CreatedByRef:       to.StringPtr("<created-by-ref>"),
				DisplayName:        to.StringPtr("<display-name>"),
				ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{},
				GranularMarkings:   []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{},
				KillChainPhases:    []*armsecurityinsights.ThreatIntelligenceKillChainPhase{},
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
