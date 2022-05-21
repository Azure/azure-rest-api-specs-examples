Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified IpAllocation.
 *
 * @summary Deletes the specified IpAllocation.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpAllocationDelete.json
 */
async function deleteIPAllocation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ipAllocationName = "test-ipallocation";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipAllocations.beginDeleteAndWait(resourceGroupName, ipAllocationName);
  console.log(result);
}

deleteIPAllocation().catch(console.error);
```
