const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Virtual Router.
 *
 * @summary Gets the specified Virtual Router.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterGet.json
 */
async function getVirtualRouter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualRouterName = "virtualRouter";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualRouters.get(resourceGroupName, virtualRouterName);
  console.log(result);
}

getVirtualRouter().catch(console.error);
