Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual machine image in an edge zone.
 *
 * @summary Gets a virtual machine image in an edge zone.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MinimumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneGetMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const offer = "aaaaaaaaaaa";
  const skus = "aaaaaaaaaaaaaaaaaa";
  const version = "aa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.get(
    location,
    edgeZone,
    publisherName,
    offer,
    skus,
    version
  );
  console.log(result);
}

virtualMachineImagesEdgeZoneGetMinimumSetGen().catch(console.error);
```
