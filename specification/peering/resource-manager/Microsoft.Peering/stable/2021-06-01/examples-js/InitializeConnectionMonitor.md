Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function initializePeeringServiceForConnectionMonitorFunctionality() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.initializeConnectionMonitor();
  console.log(result);
}

initializePeeringServiceForConnectionMonitorFunctionality().catch(console.error);
```
