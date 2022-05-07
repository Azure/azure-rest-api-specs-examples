Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing long term retention backup.
 *
 * @summary Updates an existing long term retention backup.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupUpdate.json
 */
async function updateTheLongTermRetentionBackup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "japaneast";
  const longTermRetentionServerName = "testserver";
  const longTermRetentionDatabaseName = "testDatabase";
  const backupName = "55555555-6666-7777-8888-999999999999;131637960820000000";
  const parameters = {
    requestedBackupStorageRedundancy: "Geo",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.longTermRetentionBackups.beginUpdateAndWait(
    locationName,
    longTermRetentionServerName,
    longTermRetentionDatabaseName,
    backupName,
    parameters
  );
  console.log(result);
}

updateTheLongTermRetentionBackup().catch(console.error);
```
