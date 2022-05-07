Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-postgresql_6.0.1/sdk/postgresql/arm-postgresql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a PostgreSQL Server key.
 *
 * @summary Creates or updates a PostgreSQL Server key.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyCreateOrUpdate.json
 */
async function createsOrUpdatesAPostgreSqlServerKey() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const serverName = "testserver";
  const keyName = "someVault_someKey_01234567890123456789012345678901";
  const resourceGroupName = "testrg";
  const parameters = {
    serverKeyType: "AzureKeyVault",
    uri: "https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901",
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.serverKeys.beginCreateOrUpdateAndWait(
    serverName,
    keyName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesAPostgreSqlServerKey().catch(console.error);
```
