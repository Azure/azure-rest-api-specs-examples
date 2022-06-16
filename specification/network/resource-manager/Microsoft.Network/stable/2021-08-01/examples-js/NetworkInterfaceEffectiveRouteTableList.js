const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all route tables applied to a network interface.
 *
 * @summary Gets all route tables applied to a network interface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkInterfaceEffectiveRouteTableList.json
 */
async function showNetworkInterfaceEffectiveRouteTables() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "nic1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.beginGetEffectiveRouteTableAndWait(
    resourceGroupName,
    networkInterfaceName
  );
  console.log(result);
}

showNetworkInterfaceEffectiveRouteTables().catch(console.error);
