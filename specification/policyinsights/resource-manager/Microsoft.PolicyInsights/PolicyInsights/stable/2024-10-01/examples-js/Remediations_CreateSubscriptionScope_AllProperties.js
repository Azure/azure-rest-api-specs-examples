const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a remediation at subscription scope.
 *
 * @summary creates or updates a remediation at subscription scope.
 * x-ms-original-file: 2024-10-01/Remediations_CreateSubscriptionScope_AllProperties.json
 */
async function createRemediationAtSubscriptionScopeWithAllProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.remediations.createOrUpdateAtSubscription("storageRemediation", {
    failureThreshold: { percentage: 0.1 },
    filters: { locations: ["eastus", "westus"] },
    parallelDeployments: 6,
    policyAssignmentId:
      "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
    policyDefinitionReferenceId: "8c8fa9e4",
    resourceCount: 42,
    resourceDiscoveryMode: "ReEvaluateCompliance",
  });
  console.log(result);
}
