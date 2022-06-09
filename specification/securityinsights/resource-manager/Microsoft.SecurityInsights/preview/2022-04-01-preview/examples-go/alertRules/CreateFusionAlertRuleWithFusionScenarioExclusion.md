```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
func ExampleAlertRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<rule-id>",
		&armsecurityinsights.FusionAlertRule{
			Etag: to.Ptr("<etag>"),
			Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
			Properties: &armsecurityinsights.FusionAlertRuleProperties{
				AlertRuleTemplateName: to.Ptr("<alert-rule-template-name>"),
				Enabled:               to.Ptr(true),
				SourceSettings: []*armsecurityinsights.FusionSourceSettings{
					{
						Enabled:    to.Ptr(true),
						SourceName: to.Ptr("<source-name>"),
					},
					{
						Enabled:    to.Ptr(true),
						SourceName: to.Ptr("<source-name>"),
						SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							},
							{
								Enabled: to.Ptr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
										},
										{
											Enabled:  to.Ptr(true),
											Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
										}},
								},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							}},
					},
					{
						Enabled:    to.Ptr(true),
						SourceName: to.Ptr("<source-name>"),
						SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
							{
								Enabled:           to.Ptr(true),
								SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
								SourceSubTypeName: to.Ptr("<source-sub-type-name>"),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
