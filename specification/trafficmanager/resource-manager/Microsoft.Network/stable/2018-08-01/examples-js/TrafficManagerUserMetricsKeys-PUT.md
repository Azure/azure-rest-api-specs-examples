```javascript
const { TrafficManagerManagementClient } = require("@azure/arm-trafficmanager");
const { DefaultAzureCredential } = require("@azure/identity");

async function trafficManagerUserMetricsKeysPut() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new TrafficManagerManagementClient(credential, subscriptionId);
  const result = await client.trafficManagerUserMetricsKeys.createOrUpdate();
  console.log(result);
}

trafficManagerUserMetricsKeysPut().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-trafficmanager_6.0.1/sdk/trafficmanager/arm-trafficmanager/README.md) on how to add the SDK to your project and authenticate.
