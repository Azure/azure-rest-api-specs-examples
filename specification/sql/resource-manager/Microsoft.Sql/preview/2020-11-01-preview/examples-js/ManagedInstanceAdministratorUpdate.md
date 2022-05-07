Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed instance administrator.
 *
 * @summary Creates or updates a managed instance administrator.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceAdministratorUpdate.json
 */
async function updateAdministratorOfManagedInstance() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const administratorName = "ActiveDirectory";
  const parameters = {
    administratorType: "ActiveDirectory",
    login: "bob@contoso.com",
    sid: "44444444-3333-2222-1111-000000000000",
    tenantId: "55555555-4444-3333-2222-111111111111",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceAdministrators.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    administratorName,
    parameters
  );
  console.log(result);
}

updateAdministratorOfManagedInstance().catch(console.error);
```
