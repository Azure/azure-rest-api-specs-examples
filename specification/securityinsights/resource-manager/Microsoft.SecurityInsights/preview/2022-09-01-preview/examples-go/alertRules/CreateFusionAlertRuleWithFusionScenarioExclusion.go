package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "myFirstFusionRule", &armsecurityinsights.FusionAlertRule{
		Etag: to.Ptr("3d00c3ca-0000-0100-0000-5d42d5010000"),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		Properties: &armsecurityinsights.FusionAlertRuleProperties{
			AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
			Enabled:               to.Ptr(true),
			SourceSettings: []*armsecurityinsights.FusionSourceSettings{
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Anomalies"),
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Alert providers"),
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
							SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
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
							SourceSubTypeName: to.Ptr("Azure Defender"),
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
							SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
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
							SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
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
							SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
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
							SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
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
							SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
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
							SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
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
							SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
						}},
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Raw logs from other sources"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled:           to.Ptr(true),
							SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
							SourceSubTypeName: to.Ptr("Palo Alto Networks"),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
