const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
 *
 * @summary Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2019-10-17-preview/examples/PrivateLinkScopePrivateLinkResourceGet.json
 */
async function getsPrivateEndpointConnection() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "MyResourceGroup";
  const scopeName = "MyPrivateLinkScope";
  const groupName = "azuremonitor";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, scopeName, groupName);
  console.log(result);
}

getsPrivateEndpointConnection().catch(console.error);
