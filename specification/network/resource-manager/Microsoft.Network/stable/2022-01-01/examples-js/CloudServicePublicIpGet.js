const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified public IP address in a cloud service.
 *
 * @summary Get the specified public IP address in a cloud service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/CloudServicePublicIpGet.json
 */
async function getVmssPublicIP() {
  const subscriptionId = "subid";
  const resourceGroupName = "cs-tester";
  const cloudServiceName = "cs1";
  const roleInstanceName = "Test_VM_0";
  const networkInterfaceName = "nic1";
  const ipConfigurationName = "ip1";
  const publicIpAddressName = "pub1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.getCloudServicePublicIPAddress(
    resourceGroupName,
    cloudServiceName,
    roleInstanceName,
    networkInterfaceName,
    ipConfigurationName,
    publicIpAddressName
  );
  console.log(result);
}

getVmssPublicIP().catch(console.error);
