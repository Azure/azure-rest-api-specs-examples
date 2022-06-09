```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the quantity of available pre-provisioned Event Hubs Clusters, indexed by Azure region.
 *
 * @summary List the quantity of available pre-provisioned Event Hubs Clusters, indexed by Azure region.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/Clusters/ListAvailableClustersGet.json
 */
async function listAvailableClusters() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.clusters.listAvailableClusterRegion();
  console.log(result);
}

listAvailableClusters().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.
