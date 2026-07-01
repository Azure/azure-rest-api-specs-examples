const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a remediation at subscription scope.
 *
 * @summary creates or updates a remediation at subscription scope.
 * x-ms-original-file: 2024-10-01/Remediations_CreateSubscriptionScope.json
 */
async function createRemediationAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.remediations.createOrUpdateAtSubscription("storageRemediation", {
    policyAssignmentId:
      "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
  });
  console.log(result);
}
