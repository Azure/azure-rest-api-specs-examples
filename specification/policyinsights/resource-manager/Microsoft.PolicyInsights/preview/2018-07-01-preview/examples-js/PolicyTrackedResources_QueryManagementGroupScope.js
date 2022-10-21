const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy tracked resources under the management group.
 *
 * @summary Queries policy tracked resources under the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScope.json
 */
async function queryAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupName = "myManagementGroup";
  const policyTrackedResourcesResource = "default";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyTrackedResources.listQueryResultsForManagementGroup(
    managementGroupName,
    policyTrackedResourcesResource
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryAtManagementGroupScope().catch(console.error);
