const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Azure Monitor PrivateLinkScope.
 *
 * @summary Deletes a Azure Monitor PrivateLinkScope.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2019-10-17-preview/examples/PrivateLinkScopesDelete.json
 */
async function privateLinkScopesDelete() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.beginDeleteAndWait(resourceGroupName, scopeName);
  console.log(result);
}

privateLinkScopesDelete().catch(console.error);
