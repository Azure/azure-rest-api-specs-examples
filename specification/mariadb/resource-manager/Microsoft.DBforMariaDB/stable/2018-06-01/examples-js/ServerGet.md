Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mariadb_2.0.1/sdk/mariadb/arm-mariadb/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a server.
 *
 * @summary Gets information about a server.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerGet.json
 */
async function serverGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mariadbtestsvc4";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.servers.get(resourceGroupName, serverName);
  console.log(result);
}

serverGet().catch(console.error);
```
