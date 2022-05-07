Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List recommendation action advisors.
 *
 * @summary List recommendation action advisors.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/AdvisorsListByServer.json
 */
async function advisorsListByServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.advisors.listByServer(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

advisorsListByServer().catch(console.error);
```
