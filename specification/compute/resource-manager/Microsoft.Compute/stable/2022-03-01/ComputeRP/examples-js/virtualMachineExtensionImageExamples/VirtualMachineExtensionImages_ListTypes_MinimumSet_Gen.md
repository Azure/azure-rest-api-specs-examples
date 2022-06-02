Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine extension image types.
 *
 * @summary Gets a list of virtual machine extension image types.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_ListTypes_MinimumSet_Gen.json
 */
async function virtualMachineExtensionImagesListTypesMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaa";
  const publisherName = "aa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listTypes(location, publisherName);
  console.log(result);
}

virtualMachineExtensionImagesListTypesMinimumSetGen().catch(console.error);
```
