const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List network managers in a resource group.
 *
 * @summary List network managers in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerList.json
 */
async function listNetworkManager() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkManagers.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkManager().catch(console.error);
