Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual machine image.
 *
 * @summary Gets a virtual machine image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImages_Get_MinimumSet_Gen.json
 */
async function virtualMachineImagesGetMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaa";
  const offer = "aa";
  const skus = "aaaaaaaaa";
  const version = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.get(
    location,
    publisherName,
    offer,
    skus,
    version
  );
  console.log(result);
}

virtualMachineImagesGetMinimumSetGen().catch(console.error);
```
