Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrades one or more virtual machines to the latest SKU set in the VM scale set model.
 *
 * @summary Upgrades one or more virtual machines to the latest SKU set in the VM scale set model.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_UpdateInstances_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetsUpdateInstancesMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const vmInstanceIDs = {
    instanceIds: ["aaaaaaaaaaaaaaaaaaaaaaaaa"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginUpdateInstancesAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmInstanceIDs
  );
  console.log(result);
}

virtualMachineScaleSetsUpdateInstancesMaximumSetGen().catch(console.error);
```
