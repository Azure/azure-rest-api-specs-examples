Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified network interface ip configuration in a virtual machine scale set.
 *
 * @summary Get the specified network interface ip configuration in a virtual machine scale set.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssNetworkInterfaceIpConfigList.json
 */
async function listVirtualMachineScaleSetNetworkInterfaceIPConfigurations() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualMachineScaleSetName = "vmss1";
  const virtualmachineIndex = "2";
  const networkInterfaceName = "nic1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkInterfaces.listVirtualMachineScaleSetIpConfigurations(
    resourceGroupName,
    virtualMachineScaleSetName,
    virtualmachineIndex,
    networkInterfaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualMachineScaleSetNetworkInterfaceIPConfigurations().catch(console.error);
```
