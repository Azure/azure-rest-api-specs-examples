Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of inaccessible databases in a logical server
 *
 * @summary Gets a list of inaccessible databases in a logical server
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ListVCoreInaccessibleDatabasesByServer.json
 */
async function getsAListOfInaccessibleDatabasesInALogicalServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databases.listInaccessibleByServer(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsAListOfInaccessibleDatabasesInALogicalServer().catch(console.error);
```
