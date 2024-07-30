const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all public IP addresses in a virtual machine IP configuration in a virtual machine scale set.
 *
 * @summary Gets information about all public IP addresses in a virtual machine IP configuration in a virtual machine scale set.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VmssVmPublicIpList.json
 */
async function listVmssvmPublicIP() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "vmss-tester";
  const virtualMachineScaleSetName = "vmss1";
  const virtualmachineIndex = "1";
  const networkInterfaceName = "nic1";
  const ipConfigurationName = "ip1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.publicIPAddresses.listVirtualMachineScaleSetVMPublicIPAddresses(
    resourceGroupName,
    virtualMachineScaleSetName,
    virtualmachineIndex,
    networkInterfaceName,
    ipConfigurationName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
