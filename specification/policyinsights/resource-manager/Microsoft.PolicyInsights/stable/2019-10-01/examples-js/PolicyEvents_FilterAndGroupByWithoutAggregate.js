const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the resources under the subscription.
 *
 * @summary Queries policy events for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndGroupByWithoutAggregate.json
 */
async function filterAndGroupWithoutAggregate() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const top = 2;
  const fromParam = new Date("2018-01-05T18:00:00Z");
  const filter = "PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append'";
  const apply =
    "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))";
  const options = {
    queryOptions: { top: top, from: fromParam, filter: filter, apply: apply },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForSubscription(
    policyEventsResource,
    subscriptionId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
