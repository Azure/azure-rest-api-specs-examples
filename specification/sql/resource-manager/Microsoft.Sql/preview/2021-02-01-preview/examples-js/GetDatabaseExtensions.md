```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database extension. This will return resource not found as it is not supported.
 *
 * @summary Gets a database extension. This will return resource not found as it is not supported.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetDatabaseExtensions.json
 */
async function getDatabaseExtensions() {
  const subscriptionId = "a3473687-7581-41e1-ac24-6bcca5843f07";
  const resourceGroupName = "rg_a1f9d6f8-30d5-4228-9504-8a364361bca3";
  const serverName = "srv_65858e0f-b1d1-4bdc-8351-a7da86ca4939";
  const databaseName = "11aa6c5e-58ed-4693-b303-3b8e3131deaa";
  const extensionName = "polybaseimport";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseExtensionsOperations.get(
    resourceGroupName,
    serverName,
    databaseName,
    extensionName
  );
  console.log(result);
}

getDatabaseExtensions().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.
