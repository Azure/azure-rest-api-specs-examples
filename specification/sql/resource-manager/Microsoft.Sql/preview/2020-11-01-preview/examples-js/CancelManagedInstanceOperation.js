const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels the asynchronous operation on the managed instance.
 *
 * @summary Cancels the asynchronous operation on the managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CancelManagedInstanceOperation.json
 */
async function cancelTheManagedInstanceManagementOperation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const operationId = "11111111-1111-1111-1111-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceOperations.cancel(
    resourceGroupName,
    managedInstanceName,
    operationId
  );
  console.log(result);
}

cancelTheManagedInstanceManagementOperation().catch(console.error);
