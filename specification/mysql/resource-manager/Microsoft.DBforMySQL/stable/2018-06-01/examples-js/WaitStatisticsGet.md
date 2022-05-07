Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve wait statistics for specified identifier.
 *
 * @summary Retrieve wait statistics for specified identifier.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsGet.json
 */
async function waitStatisticsGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const waitStatisticsId =
    "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.waitStatistics.get(resourceGroupName, serverName, waitStatisticsId);
  console.log(result);
}

waitStatisticsGet().catch(console.error);
```
