const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network interfaces in a virtual machine scale set.
 *
 * @summary Gets all network interfaces in a virtual machine scale set.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VmssNetworkInterfaceList.json
 */
async function listVirtualMachineScaleSetNetworkInterfaces() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualMachineScaleSetName = "vmss1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkInterfaces.listVirtualMachineScaleSetNetworkInterfaces(
    resourceGroupName,
    virtualMachineScaleSetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualMachineScaleSetNetworkInterfaces().catch(console.error);
