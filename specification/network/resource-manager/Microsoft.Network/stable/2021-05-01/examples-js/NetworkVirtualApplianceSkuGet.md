Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves a single available sku for network virtual appliance.
 *
 * @summary Retrieves a single available sku for network virtual appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualApplianceSkuGet.json
 */
async function networkVirtualApplianceSkuGet() {
  const subscriptionId = "subid";
  const skuName = "ciscoSdwan";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualApplianceSkus.get(skuName);
  console.log(result);
}

networkVirtualApplianceSkuGet().catch(console.error);
```
