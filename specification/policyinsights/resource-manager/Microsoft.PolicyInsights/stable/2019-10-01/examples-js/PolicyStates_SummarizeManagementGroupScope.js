const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Summarizes policy states for the resources under the management group.
 *
 * @summary Summarizes policy states for the resources under the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
 */
async function summarizeAtManagementGroupScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyStatesSummaryResource = "latest";
  const managementGroupName = "myManagementGroup";
  const top = 0;
  const fromParam = new Date("2019-10-05T18:00:00Z");
  const to = new Date("2019-10-06T18:00:00Z");
  const filter = "PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit'";
  const options = {
    queryOptions: { top: top, from: fromParam, to: to, filter: filter },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.summarizeForManagementGroup(
    policyStatesSummaryResource,
    managementGroupName,
    options
  );
  console.log(result);
}
