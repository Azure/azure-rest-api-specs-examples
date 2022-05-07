Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-mysql_5.0.1/sdk/mysql/arm-mysql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a list of configurations in a given server.
 *
 * @summary Update a list of configurations in a given server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
 */
async function configurationList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc1";
  const value = {};
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.serverParameters.beginListUpdateConfigurationsAndWait(
    resourceGroupName,
    serverName,
    value
  );
  console.log(result);
}

configurationList().catch(console.error);
```
