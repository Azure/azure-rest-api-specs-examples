const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the resources under the subscription.
 *
 * @summary Queries policy events for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndMultipleGroups.json
 */
async function filterAndMultipleGroups() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const top = 10;
  const orderBy = "NumDeniedResources desc";
  const fromParam = new Date("2018-01-01T00:00:00Z");
  const filter = "PolicyDefinitionAction eq 'deny'";
  const apply =
    "groupby((PolicyAssignmentId, PolicyDefinitionId, ResourceId))/groupby((PolicyAssignmentId, PolicyDefinitionId), aggregate($count as NumDeniedResources))";
  const options = {
    queryOptions: { top: top, orderBy: orderBy, from: fromParam, filter: filter, apply: apply },
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

filterAndMultipleGroups().catch(console.error);
