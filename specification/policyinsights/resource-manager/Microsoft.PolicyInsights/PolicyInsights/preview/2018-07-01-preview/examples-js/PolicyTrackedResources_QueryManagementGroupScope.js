const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy tracked resources under the management group.
 *
 * @summary queries policy tracked resources under the management group.
 * x-ms-original-file: 2018-07-01-preview/PolicyTrackedResources_QueryManagementGroupScope.json
 */
async function queryAtManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyTrackedResources.listQueryResultsForManagementGroup(
    "myManagementGroup",
    "default",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
