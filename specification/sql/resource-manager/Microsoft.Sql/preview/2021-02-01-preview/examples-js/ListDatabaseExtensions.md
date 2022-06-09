```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List database extension. This will return an empty list as it is not supported.
 *
 * @summary List database extension. This will return an empty list as it is not supported.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListDatabaseExtensions.json
 */
async function listDatabaseExtensions() {
  const subscriptionId = "7b2515fe-f230-4017-8cf0-695163acab85";
  const resourceGroupName = "rg_4007c5a9-b3b0-41e1-bd46-9eef38768a4a";
  const serverName = "srv_3b67ec2a-519b-43a7-8533-fb62dce3431e";
  const databaseName = "719d8fa4-bf0f-48fc-8cd3-ef40fe6ba1fe";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.databaseExtensionsOperations.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDatabaseExtensions().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.
