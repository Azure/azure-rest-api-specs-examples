const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resources under the subscription.
 *
 * @summary queries policy events for the resources under the subscription.
 * x-ms-original-file: 2024-10-01/PolicyEvents_FilterAndGroupByWithoutAggregate.json
 */
async function filterAndGroupWithoutAggregate() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForSubscription(
    "default",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    {
      queryOptions: {
        top: 2,
        from: new Date("2018-01-05T18:00:00Z"),
        filter: "PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append'",
        apply:
          "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))",
      },
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
