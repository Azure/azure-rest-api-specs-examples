const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified network interface.
 *
 * @summary Deletes the specified network interface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkInterfaceDelete.json
 */
async function deleteNetworkInterface() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "test-nic";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.beginDeleteAndWait(
    resourceGroupName,
    networkInterfaceName
  );
  console.log(result);
}

deleteNetworkInterface().catch(console.error);
