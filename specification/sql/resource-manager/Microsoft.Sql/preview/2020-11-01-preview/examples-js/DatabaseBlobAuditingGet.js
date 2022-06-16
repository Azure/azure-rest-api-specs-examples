const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database's blob auditing policy.
 *
 * @summary Gets a database's blob auditing policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseBlobAuditingGet.json
 */
async function getADatabaseBlobAuditingPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-6852";
  const serverName = "blobauditingtest-2080";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseBlobAuditingPolicies.get(
    resourceGroupName,
    serverName,
    databaseName
  );
  console.log(result);
}

getADatabaseBlobAuditingPolicy().catch(console.error);
