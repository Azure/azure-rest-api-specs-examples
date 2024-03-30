const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private endpoint connections on a private link scope.
 *
 * @summary Gets all private endpoint connections on a private link scope.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceList.json
 */
async function getsListOfScopedResourcesInAPrivateLinkScope() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "MyResourceGroup";
  const scopeName = "MyPrivateLinkScope";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkScopedResources.listByPrivateLinkScope(
    resourceGroupName,
    scopeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
