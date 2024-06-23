const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all public IP addresses on a virtual machine scale set level.
 *
 * @summary Gets information about all public IP addresses on a virtual machine scale set level.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VmssPublicIpListAll.json
 */
async function listVmssPublicIP() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "vmss-tester";
  const virtualMachineScaleSetName = "vmss1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.publicIPAddresses.listVirtualMachineScaleSetPublicIPAddresses(
    resourceGroupName,
    virtualMachineScaleSetName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
