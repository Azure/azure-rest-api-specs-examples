Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 *
 * @summary Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/EffectiveRoutesListForVirtualHub.json
 */
async function effectiveRoutesForTheVirtualHub() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const effectiveRoutesParameters = {};
  const options = {
    effectiveRoutesParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.beginGetEffectiveVirtualHubRoutesAndWait(
    resourceGroupName,
    virtualHubName,
    options
  );
  console.log(result);
}

effectiveRoutesForTheVirtualHub().catch(console.error);
```
