Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a disk access resource.
 *
 * @summary Updates (patches) a disk access resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/UpdateADiskAccess.json
 */
async function updateADiskAccessResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const diskAccess = {
    tags: { department: "Development", project: "PrivateEndpoints" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginUpdateAndWait(
    resourceGroupName,
    diskAccessName,
    diskAccess
  );
  console.log(result);
}

updateADiskAccessResource().catch(console.error);
```
