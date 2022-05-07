Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a database's transparent data encryption configuration.
 *
 * @summary Updates a database's transparent data encryption configuration.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedTransparentDataEncryptionUpdate.json
 */
async function updateADatabaseTransparentDataEncryptionStateWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securitytde-42-rg";
  const managedInstanceName = "securitytde-42";
  const databaseName = "testdb";
  const tdeName = "current";
  const parameters = { state: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseTransparentDataEncryption.createOrUpdate(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    tdeName,
    parameters
  );
  console.log(result);
}

updateADatabaseTransparentDataEncryptionStateWithMinimalParameters().catch(console.error);
```
