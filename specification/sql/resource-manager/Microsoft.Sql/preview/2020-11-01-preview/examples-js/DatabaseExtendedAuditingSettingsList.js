const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists extended auditing settings of a database.
 *
 * @summary Lists extended auditing settings of a database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseExtendedAuditingSettingsList.json
 */
async function listExtendedAuditingSettingsOfADatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-6852";
  const serverName = "blobauditingtest-2080";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extendedDatabaseBlobAuditingPolicies.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExtendedAuditingSettingsOfADatabase().catch(console.error);
