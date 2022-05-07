Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed database's transparent data encryption.
 *
 * @summary Gets a managed database's transparent data encryption.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedTransparentDataEncryptionGet.json
 */
async function getADatabaseTransparentDataEncryption() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "security-tde-resourcegroup";
  const managedInstanceName = "securitytde";
  const databaseName = "testdb";
  const tdeName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseTransparentDataEncryption.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    tdeName
  );
  console.log(result);
}

getADatabaseTransparentDataEncryption().catch(console.error);
```
