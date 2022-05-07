Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all virtualClusters in the subscription.
 *
 * @summary Gets a list of all virtualClusters in the subscription.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterList.json
 */
async function listVirtualClusters() {
  const subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualClusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualClusters().catch(console.error);
```
