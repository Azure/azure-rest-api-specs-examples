const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of security events.
 *
 * @summary Gets a list of security events.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityEventsGetMin.json
 */
async function getTheManagedDatabaseSecurityEventsWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testcl";
  const databaseName = "database1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseSecurityEvents.listByDatabase(
    resourceGroupName,
    managedInstanceName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheManagedDatabaseSecurityEventsWithMinimalParameters().catch(console.error);
