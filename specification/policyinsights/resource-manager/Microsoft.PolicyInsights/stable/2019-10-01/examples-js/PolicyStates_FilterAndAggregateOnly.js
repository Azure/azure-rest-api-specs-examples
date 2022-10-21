const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the resources under the subscription.
 *
 * @summary Queries policy states for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndAggregateOnly.json
 */
async function filterAndAggregateOnly() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "latest";
  const fromParam = new Date("2019-10-05T18:00:00Z");
  const filter = "PolicyDefinitionAction eq 'deny'";
  const apply = "aggregate($count as NumDenyStates)";
  const options = {
    queryOptions: { from: fromParam, filter: filter, apply: apply },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForSubscription(
    policyStatesResource,
    subscriptionId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

filterAndAggregateOnly().catch(console.error);
