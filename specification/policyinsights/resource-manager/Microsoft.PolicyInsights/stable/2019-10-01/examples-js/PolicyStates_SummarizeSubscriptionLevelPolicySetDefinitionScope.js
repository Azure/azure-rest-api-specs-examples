const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Summarizes policy states for the subscription level policy set definition.
 *
 * @summary Summarizes policy states for the subscription level policy set definition.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicySetDefinitionScope.json
 */
async function summarizeAtPolicySetDefinitionScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesSummaryResource = "latest";
  const policySetDefinitionName = "3e3807c1-65c9-49e0-a406-82d8ae3e338c";
  const top = 1;
  const fromParam = new Date("2019-10-05T18:00:00Z");
  const to = new Date("2019-10-06T18:00:00Z");
  const filter = "PolicyDefinitionAction eq 'deny'";
  const options = {
    queryOptions: { top: top, to: to, from: fromParam, filter: filter },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.summarizeForPolicySetDefinition(
    policyStatesSummaryResource,
    subscriptionId,
    policySetDefinitionName,
    options
  );
  console.log(result);
}

summarizeAtPolicySetDefinitionScope().catch(console.error);
