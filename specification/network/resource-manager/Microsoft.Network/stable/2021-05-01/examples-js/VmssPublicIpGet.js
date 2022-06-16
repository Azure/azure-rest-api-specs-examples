const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified public IP address in a virtual machine scale set.
 *
 * @summary Get the specified public IP address in a virtual machine scale set.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssPublicIpGet.json
 */
async function getVmssPublicIP() {
  const subscriptionId = "subid";
  const resourceGroupName = "vmss-tester";
  const virtualMachineScaleSetName = "vmss1";
  const virtualmachineIndex = "1";
  const networkInterfaceName = "nic1";
  const ipConfigurationName = "ip1";
  const publicIpAddressName = "pub1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.getVirtualMachineScaleSetPublicIPAddress(
    resourceGroupName,
    virtualMachineScaleSetName,
    virtualmachineIndex,
    networkInterfaceName,
    ipConfigurationName,
    publicIpAddressName
  );
  console.log(result);
}

getVmssPublicIP().catch(console.error);
