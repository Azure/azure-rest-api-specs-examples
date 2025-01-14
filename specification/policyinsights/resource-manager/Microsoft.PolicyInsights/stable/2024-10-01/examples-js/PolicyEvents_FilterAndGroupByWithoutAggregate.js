const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy events for the resources under the subscription.
 *
 * @summary Queries policy events for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyEvents_FilterAndGroupByWithoutAggregate.json
 */
async function filterAndGroupWithoutAggregate() {
  const policyEventsResource = "default";
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const top = 2;
  const fromParam = new Date("2018-01-05T18:00:00Z");
  const filter = "PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append'";
  const apply =
    "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))";
  const options = {
    top,
    fromParam,
    filter,
    apply,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForSubscription(
    policyEventsResource,
    subscriptionId,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
