Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetExtensionsGetMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const vmssExtensionName = "aaaaaaaaaaaaaaaaaaaa";
  const expand = "aaaaaaa";
  const options = { expand: expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetExtensions.get(
    resourceGroupName,
    vmScaleSetName,
    vmssExtensionName,
    options
  );
  console.log(result);
}

virtualMachineScaleSetExtensionsGetMaximumSetGen().catch(console.error);
```
