```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List network usages for a subscription.
 *
 * @summary List network usages for a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/UsageListSpacedLocation.json
 */
async function listUsagesSpacedLocation() {
  const subscriptionId = "subid";
  const location = "West US";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listUsagesSpacedLocation().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
