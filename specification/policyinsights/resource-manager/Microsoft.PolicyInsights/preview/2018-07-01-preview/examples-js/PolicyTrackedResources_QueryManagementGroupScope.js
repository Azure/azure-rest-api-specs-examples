const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy tracked resources under the management group.
 *
 * @summary Queries policy tracked resources under the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScope.json
 */
async function queryAtManagementGroupScope() {
  const managementGroupName = "myManagementGroup";
  const policyTrackedResourcesResource = "default";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyTrackedResources.listQueryResultsForManagementGroup(
    managementGroupName,
    policyTrackedResourcesResource,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
