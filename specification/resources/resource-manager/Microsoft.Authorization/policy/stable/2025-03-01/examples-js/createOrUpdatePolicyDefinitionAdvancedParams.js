const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy definition in the given subscription with the given name.
 *
 * @summary this operation creates or updates a policy definition in the given subscription with the given name.
 * x-ms-original-file: 2025-03-01/createOrUpdatePolicyDefinitionAdvancedParams.json
 */
async function createOrUpdateAPolicyDefinitionWithAdvancedParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.createOrUpdate("EventHubDiagnosticLogs", {
    description:
      "Audit enabling of logs and retain them up to a year. This enables recreation of activity trails for investigation purposes when a security incident occurs or your network is compromised",
    displayName: "Event Hubs should have diagnostic logging enabled",
    metadata: { category: "Event Hub" },
    mode: "Indexed",
    parameters: {
      requiredRetentionDays: {
        type: "Integer",
        allowedValues: [0, 30, 90, 180, 365],
        defaultValue: 365,
        metadata: {
          description: "The required diagnostic logs retention in days",
          displayName: "Required retention (days)",
        },
      },
    },
    policyRule: {
      if: { equals: "Microsoft.EventHub/namespaces", field: "type" },
      then: {
        effect: "AuditIfNotExists",
        details: {
          type: "Microsoft.Insights/diagnosticSettings",
          existenceCondition: {
            allOf: [
              {
                equals: "true",
                field: "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled",
              },
              {
                equals: "[parameters('requiredRetentionDays')]",
                field: "Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days",
              },
            ],
          },
        },
      },
    },
  });
  console.log(result);
}
