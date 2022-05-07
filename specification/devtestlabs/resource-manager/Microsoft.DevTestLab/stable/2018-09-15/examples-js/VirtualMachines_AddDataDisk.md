Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Attach a new or existing data disk to virtual machine. This operation can take a while to complete.
 *
 * @summary Attach a new or existing data disk to virtual machine. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_AddDataDisk.json
 */
async function virtualMachinesAddDataDisk() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{virtualMachineName}";
  const dataDiskProperties = {
    attachNewDataDiskOptions: {
      diskName: "{diskName}",
      diskSizeGiB: 127,
      diskType: "{diskType}",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginAddDataDiskAndWait(
    resourceGroupName,
    labName,
    name,
    dataDiskProperties
  );
  console.log(result);
}

virtualMachinesAddDataDisk().catch(console.error);
```
