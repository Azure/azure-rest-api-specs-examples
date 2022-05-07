Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Service Fabric cluster resources created or in the process of being created in the subscription.
 *
 * @summary Gets all Service Fabric cluster resources created or in the process of being created in the subscription.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListOperation_example.json
 */
async function listClusters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.clusters.list();
  console.log(result);
}

listClusters().catch(console.error);
```
