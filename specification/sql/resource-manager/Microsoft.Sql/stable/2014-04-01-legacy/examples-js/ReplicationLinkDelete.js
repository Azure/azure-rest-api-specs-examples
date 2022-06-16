const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a database replication link. Cannot be done during failover.
 *
 * @summary Deletes a database replication link. Cannot be done during failover.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ReplicationLinkDelete.json
 */
async function deleteAReplicationLink() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-4799";
  const serverName = "sqlcrudtest-6440";
  const databaseName = "testdb";
  const linkId = "5b301b68-03f6-4b26-b0f4-73ebb8634238";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.replicationLinks.delete(
    resourceGroupName,
    serverName,
    databaseName,
    linkId
  );
  console.log(result);
}

deleteAReplicationLink().catch(console.error);
