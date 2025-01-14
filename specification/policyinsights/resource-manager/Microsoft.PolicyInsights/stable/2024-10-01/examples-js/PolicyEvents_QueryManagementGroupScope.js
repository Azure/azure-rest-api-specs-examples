const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy events for the resources under the management group.
 *
 * @summary Queries policy events for the resources under the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyEvents_QueryManagementGroupScope.json
 */
async function queryAtManagementGroupScope() {
  const policyEventsResource = "default";
  const managementGroupName = "myManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForManagementGroup(
    policyEventsResource,
    managementGroupName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
