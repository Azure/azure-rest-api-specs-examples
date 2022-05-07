Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-trafficmanager_6.0.1/sdk/trafficmanager/arm-trafficmanager/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Traffic Manager profiles within a subscription.
 *
 * @summary Lists all Traffic Manager profiles within a subscription.
 * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-GET-BySubscription.json
 */
async function listBySubscription() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.profiles.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBySubscription().catch(console.error);
```
