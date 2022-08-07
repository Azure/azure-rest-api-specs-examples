const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of all HubVirtualNetworkConnections.
 *
 * @summary Retrieves the details of all HubVirtualNetworkConnections.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/HubVirtualNetworkConnectionList.json
 */
async function hubVirtualNetworkConnectionList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hubVirtualNetworkConnections.list(
    resourceGroupName,
    virtualHubName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

hubVirtualNetworkConnectionList().catch(console.error);
