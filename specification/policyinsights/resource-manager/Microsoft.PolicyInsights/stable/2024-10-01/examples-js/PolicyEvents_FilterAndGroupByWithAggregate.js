const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy events for the resources under the subscription.
 *
 * @summary Queries policy events for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyEvents_FilterAndGroupByWithAggregate.json
 */
async function filterAndGroupWithAggregate() {
  const policyEventsResource = "default";
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const top = 2;
  const fromParam = new Date("2018-02-05T18:00:00Z");
  const filter = "PolicyDefinitionAction eq 'audit' or PolicyDefinitionAction eq 'deny'";
  const apply =
    "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId), aggregate($count as NumEvents))";
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
