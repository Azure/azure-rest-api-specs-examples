package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: 2025-03-01/createOrUpdatePolicyDefinitionAdvancedParams.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinitionWithAdvancedParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdate(ctx, "EventHubDiagnosticLogs", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised"),
			DisplayName: to.Ptr("Event Hubs should have diagnostic logging enabled"),
			Metadata: map[string]any{
				"category": "Event Hub",
			},
			Mode: to.Ptr("Indexed"),
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"requiredRetentionDays": {
					Type: to.Ptr(armpolicy.ParameterTypeInteger),
					AllowedValues: []any{
						0,
						30,
						90,
						180,
						365,
					},
					DefaultValue: 365,
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("The required diagnostic logs retention in days"),
						DisplayName: to.Ptr("Required retention (days)"),
					},
				},
			},
			PolicyRule: map[string]any{
				"if": map[string]any{
					"equals": "Microsoft.EventHub/namespaces",
					"field":  "type",
				},
				"then": map[string]any{
					"effect": "AuditIfNotExists",
					"details": map[string]any{
						"type": "Microsoft.Insights/diagnosticSettings",
						"existenceCondition": map[string]any{
							"allOf": []any{
								map[string]any{
									"equals": "true",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
								},
								map[string]any{
									"equals": "[parameters('requiredRetentionDays')]",
									"field":  "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
								},
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
