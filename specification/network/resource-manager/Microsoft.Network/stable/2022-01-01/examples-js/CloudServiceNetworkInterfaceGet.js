const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified network interface in a cloud service.
 *
 * @summary Get the specified network interface in a cloud service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/CloudServiceNetworkInterfaceGet.json
 */
async function getCloudServiceNetworkInterface() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const cloudServiceName = "cs1";
  const roleInstanceName = "TestVMRole_IN_0";
  const networkInterfaceName = "nic1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.getCloudServiceNetworkInterface(
    resourceGroupName,
    cloudServiceName,
    roleInstanceName,
    networkInterfaceName
  );
  console.log(result);
}

getCloudServiceNetworkInterface().catch(console.error);
