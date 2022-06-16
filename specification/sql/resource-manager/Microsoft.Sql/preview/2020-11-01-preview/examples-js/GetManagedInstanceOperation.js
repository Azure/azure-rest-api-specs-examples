const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a management operation on a managed instance.
 *
 * @summary Gets a management operation on a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetManagedInstanceOperation.json
 */
async function getsTheManagedInstanceManagementOperation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const operationId = "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceOperations.get(
    resourceGroupName,
    managedInstanceName,
    operationId
  );
  console.log(result);
}

getsTheManagedInstanceManagementOperation().catch(console.error);
