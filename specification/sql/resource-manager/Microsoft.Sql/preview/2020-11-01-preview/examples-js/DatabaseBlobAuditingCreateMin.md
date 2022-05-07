Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a database's blob auditing policy.
 *
 * @summary Creates or updates a database's blob auditing policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseBlobAuditingCreateMin.json
 */
async function createOrUpdateADatabaseBlobAuditingPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-4799";
  const serverName = "blobauditingtest-6440";
  const databaseName = "testdb";
  const parameters = {
    state: "Enabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseBlobAuditingPolicies.createOrUpdate(
    resourceGroupName,
    serverName,
    databaseName,
    parameters
  );
  console.log(result);
}

createOrUpdateADatabaseBlobAuditingPolicyWithMinimalParameters().catch(console.error);
```
