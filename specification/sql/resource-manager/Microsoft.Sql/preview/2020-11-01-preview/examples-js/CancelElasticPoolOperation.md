Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels the asynchronous operation on the elastic pool.
 *
 * @summary Cancels the asynchronous operation on the elastic pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CancelElasticPoolOperation.json
 */
async function cancelTheElasticPoolManagementOperation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-6661";
  const elasticPoolName = "testpool";
  const operationId = "f779414b-e748-4925-8cfe-c8598f7660ae";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.elasticPoolOperations.cancel(
    resourceGroupName,
    serverName,
    elasticPoolName,
    operationId
  );
  console.log(result);
}

cancelTheElasticPoolManagementOperation().catch(console.error);
```
