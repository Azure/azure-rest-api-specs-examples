Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the current ledger digest upload configuration for a database.
 *
 * @summary Gets the current ledger digest upload configuration for a database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/LedgerDigestUploadsGet.json
 */
async function getsTheCurrentLedgerDigestUploadConfigurationForADatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "ledgertestrg";
  const serverName = "ledgertestserver";
  const databaseName = "testdb";
  const ledgerDigestUploads = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.ledgerDigestUploadsOperations.get(
    resourceGroupName,
    serverName,
    databaseName,
    ledgerDigestUploads
  );
  console.log(result);
}

getsTheCurrentLedgerDigestUploadConfigurationForADatabase().catch(console.error);
```
