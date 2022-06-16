const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of all VirtualHubBgpConnections.
 *
 * @summary Retrieves the details of all VirtualHubBgpConnections.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubBgpConnectionList.json
 */
async function virtualHubRouteTableV2List() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "hub1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualHubBgpConnections.list(resourceGroupName, virtualHubName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualHubRouteTableV2List().catch(console.error);
