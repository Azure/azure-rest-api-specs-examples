package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAdvancedParams.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinitionWithAdvancedParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
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
						float64(0),
						float64(30),
						float64(90),
						float64(180),
						float64(365)},
					DefaultValue: float64(365),
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
