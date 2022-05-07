Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetExtensionsCreateOrUpdateMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaa";
  const vmssExtensionName = "aaaaaaaaaaa";
  const extensionParameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName,
    extensionParameters
  );
  console.log(result);
}

virtualMachineScaleSetExtensionsCreateOrUpdateMinimumSetGen().catch(console.error);
```
