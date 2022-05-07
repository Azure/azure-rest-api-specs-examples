Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Transfers all data disks attached to the virtual machine to be owned by the current user. This operation can take a while to complete.
 *
 * @summary Transfers all data disks attached to the virtual machine to be owned by the current user. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_TransferDisks.json
 */
async function virtualMachinesTransferDisks() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{virtualmachineName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginTransferDisksAndWait(
    resourceGroupName,
    labName,
    name
  );
  console.log(result);
}

virtualMachinesTransferDisks().catch(console.error);
```
