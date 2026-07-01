const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resources under the management group.
 *
 * @summary queries policy events for the resources under the management group.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QueryManagementGroupScope.json
 */
async function queryAtManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForManagementGroup(
    "default",
    "myManagementGroup",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
