const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed instance administrator.
 *
 * @summary Gets a managed instance administrator.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceAdministratorGet.json
 */
async function getAdministratorOfManagedInstance() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const administratorName = "ActiveDirectory";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceAdministrators.get(
    resourceGroupName,
    managedInstanceName,
    administratorName
  );
  console.log(result);
}

getAdministratorOfManagedInstance().catch(console.error);
