Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the subscription is visible to private link service in the specified resource group.
 *
 * @summary Checks whether the subscription is visible to private link service in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CheckPrivateLinkServiceVisibilityByResourceGroup.json
 */
async function checkPrivateLinkServiceVisibility() {
  const subscriptionId = "subid";
  const location = "westus";
  const resourceGroupName = "rg1";
  const parameters = {
    privateLinkServiceAlias: "mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result =
    await client.privateLinkServices.beginCheckPrivateLinkServiceVisibilityByResourceGroupAndWait(
      location,
      resourceGroupName,
      parameters
    );
  console.log(result);
}

checkPrivateLinkServiceVisibility().catch(console.error);
```
