Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of  Server keys.
 *
 * @summary Gets a list of  Server keys.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyList.json
 */
async function listTheKeysForAMySqlServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "testserver";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverKeys.list(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheKeysForAMySqlServer().catch(console.error);
```
