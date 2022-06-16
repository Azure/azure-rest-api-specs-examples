const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified network interface in a virtual machine scale set.
 *
 * @summary Get the specified network interface in a virtual machine scale set.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssNetworkInterfaceGet.json
 */
async function getVirtualMachineScaleSetNetworkInterface() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualMachineScaleSetName = "vmss1";
  const virtualmachineIndex = "1";
  const networkInterfaceName = "nic1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.getVirtualMachineScaleSetNetworkInterface(
    resourceGroupName,
    virtualMachineScaleSetName,
    virtualmachineIndex,
    networkInterfaceName
  );
  console.log(result);
}

getVirtualMachineScaleSetNetworkInterface().catch(console.error);
