package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
func ExampleAlertRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewAlertRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<rule-id>",
		&armsecurityinsights.FusionAlertRule{
			Etag: to.StringPtr("<etag>"),
			Kind: armsecurityinsights.AlertRuleKind("Fusion").ToPtr(),
			Properties: &armsecurityinsights.FusionAlertRuleProperties{
				AlertRuleTemplateName: to.StringPtr("<alert-rule-template-name>"),
				Enabled:               to.BoolPtr(true),
				SourceSettings: []*armsecurityinsights.FusionSourceSettings{
					{
						Enabled:    to.BoolPtr(true),
						SourceName: to.StringPtr("<source-name>"),
					},
					{
						Enabled:    to.BoolPtr(true),
						SourceName: to.StringPtr("<source-name>"),
						SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							},
							{
								Enabled: to.BoolPtr(true),
								SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
									Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("High").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Medium").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Low").ToPtr(),
										},
										{
											Enabled:  to.BoolPtr(true),
											Severity: armsecurityinsights.AlertSeverity("Informational").ToPtr(),
										}},
								},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							}},
					},
					{
						Enabled:    to.BoolPtr(true),
						SourceName: to.StringPtr("<source-name>"),
						SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
							{
								Enabled:           to.BoolPtr(true),
								SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
								SourceSubTypeName: to.StringPtr("<source-sub-type-name>"),
							}},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AlertRulesClientCreateOrUpdateResult)
}
