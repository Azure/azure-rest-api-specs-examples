const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed instance keys.
 *
 * @summary Gets a list of managed instance keys.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceKeyList.json
 */
async function listTheKeysForAManagedInstance() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstanceKeys.listByInstance(
    resourceGroupName,
    managedInstanceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheKeysForAManagedInstance().catch(console.error);
