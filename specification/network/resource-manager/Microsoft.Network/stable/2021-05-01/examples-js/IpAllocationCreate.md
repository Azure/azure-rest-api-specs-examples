Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an IpAllocation in the specified resource group.
 *
 * @summary Creates or updates an IpAllocation in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/IpAllocationCreate.json
 */
async function createIPAllocation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ipAllocationName = "test-ipallocation";
  const parameters = {
    typePropertiesType: "Hypernet",
    allocationTags: {
      vNetID:
        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1",
    },
    location: "centraluseuap",
    prefix: "3.2.5.0/24",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipAllocations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    ipAllocationName,
    parameters
  );
  console.log(result);
}

createIPAllocation().catch(console.error);
```
