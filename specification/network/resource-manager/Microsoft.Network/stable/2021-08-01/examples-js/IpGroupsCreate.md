Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an ipGroups in a specified resource group.
 *
 * @summary Creates or updates an ipGroups in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/IpGroupsCreate.json
 */
async function createOrUpdateIPGroups() {
  const subscriptionId = "subId";
  const resourceGroupName = "myResourceGroup";
  const ipGroupsName = "ipGroups1";
  const parameters = {
    ipAddresses: ["13.64.39.16/32", "40.74.146.80/31", "40.74.147.32/28"],
    location: "West US",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    ipGroupsName,
    parameters
  );
  console.log(result);
}

createOrUpdateIPGroups().catch(console.error);
```
