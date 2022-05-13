Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-vmwarecloudsimple_3.0.0/sdk/vmwarecloudsimple/arm-vmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch virtual machine properties
 *
 * @summary Patch virtual machine properties
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/PatchVirtualMachine.json
 */
async function patchVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const virtualMachineName = "myVirtualMachine";
  const virtualMachineRequest = { tags: { myTag: "tagValue" } };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.virtualMachines.beginUpdateAndWait(
    resourceGroupName,
    virtualMachineName,
    virtualMachineRequest
  );
  console.log(result);
}

patchVirtualMachine().catch(console.error);
```
