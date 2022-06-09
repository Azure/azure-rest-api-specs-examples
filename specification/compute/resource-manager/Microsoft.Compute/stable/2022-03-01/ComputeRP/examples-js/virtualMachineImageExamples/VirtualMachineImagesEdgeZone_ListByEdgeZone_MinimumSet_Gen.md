Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all virtual machine image versions for the specified edge zone
 *
 * @summary Gets a list of all virtual machine image versions for the specified edge zone
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MinimumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListByEdgeZoneMinimumSetGen() {
  const subscriptionId = "5ece5940-d962-4dad-a98f-ca9ac0f021a5";
  const location = "WestUS";
  const edgeZone = "microsoftlosangeles1";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listByEdgeZone(location, edgeZone);
  console.log(result);
}

virtualMachineImagesEdgeZoneListByEdgeZoneMinimumSetGen().catch(console.error);
```
