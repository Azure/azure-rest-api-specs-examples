const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Summarizes policy states for the subscription level policy definition.
 *
 * @summary Summarizes policy states for the subscription level policy definition.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyDefinitionScope.json
 */
async function summarizeAtPolicyDefinitionScope() {
  const policyStatesSummaryResource = "latest";
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const policyDefinitionName = "24813039-7534-408a-9842-eb99f45721b1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.policyStates.summarizeForPolicyDefinition(
    policyStatesSummaryResource,
    subscriptionId,
    policyDefinitionName,
  );
  console.log(result);
}
