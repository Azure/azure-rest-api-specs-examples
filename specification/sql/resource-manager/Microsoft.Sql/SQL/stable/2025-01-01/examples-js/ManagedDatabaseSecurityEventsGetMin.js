const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of security events.
 *
 * @summary gets a list of security events.
 * x-ms-original-file: 2025-01-01/ManagedDatabaseSecurityEventsGetMin.json
 */
async function getTheManagedDatabaseSecurityEventsWithMinimalParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedDatabaseSecurityEvents.listByDatabase(
    "testrg",
    "testcl",
    "database1",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
