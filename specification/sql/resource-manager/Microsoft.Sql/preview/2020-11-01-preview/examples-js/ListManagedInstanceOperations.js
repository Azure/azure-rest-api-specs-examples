const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of operations performed on the managed instance.
 *
 * @summary Gets a list of operations performed on the managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListManagedInstanceOperations.json
 */
async function listTheManagedInstanceManagementOperations() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstanceOperations.listByManagedInstance(
    resourceGroupName,
    managedInstanceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
