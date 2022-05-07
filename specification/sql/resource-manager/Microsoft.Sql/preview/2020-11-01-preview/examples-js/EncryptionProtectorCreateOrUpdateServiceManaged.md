Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing encryption protector.
 *
 * @summary Updates an existing encryption protector.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/EncryptionProtectorCreateOrUpdateServiceManaged.json
 */
async function updateTheEncryptionProtectorToServiceManaged() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const encryptionProtectorName = "current";
  const parameters = {
    serverKeyName: "ServiceManaged",
    serverKeyType: "ServiceManaged",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.encryptionProtectors.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    encryptionProtectorName,
    parameters
  );
  console.log(result);
}

updateTheEncryptionProtectorToServiceManaged().catch(console.error);
```
