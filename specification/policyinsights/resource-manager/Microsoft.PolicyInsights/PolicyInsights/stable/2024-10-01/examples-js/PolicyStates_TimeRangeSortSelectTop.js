const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy states for the resources under the subscription.
 *
 * @summary queries policy states for the resources under the subscription.
 * x-ms-original-file: 2024-10-01/PolicyStates_TimeRangeSortSelectTop.json
 */
async function timeRangeSortSelectAndLimit() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyStates.listQueryResultsForSubscription(
    "latest",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    {
      queryOptions: {
        top: 2,
        orderBy:
          "Timestamp desc, PolicyAssignmentId asc, SubscriptionId asc, ResourceGroup asc, ResourceId",
        select:
          "Timestamp, PolicyAssignmentId, PolicyDefinitionId, SubscriptionId, ResourceGroup, ResourceId, policyDefinitionGroupNames",
        from: new Date("2019-10-05T18:00:00Z"),
        to: new Date("2019-10-06T18:00:00Z"),
      },
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
