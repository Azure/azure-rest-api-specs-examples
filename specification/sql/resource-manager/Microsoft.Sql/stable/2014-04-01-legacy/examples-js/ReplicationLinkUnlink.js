const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a database replication link in forced or friendly way.
 *
 * @summary Deletes a database replication link in forced or friendly way.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ReplicationLinkUnlink.json
 */
async function deleteReplicationLink() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-8931";
  const serverName = "sqlcrudtest-2137";
  const databaseName = "testdb";
  const linkId = "f0550bf5-07ce-4270-8e4b-71737975973a";
  const parameters = { forcedTermination: true };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.replicationLinks.beginUnlinkAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    linkId,
    parameters
  );
  console.log(result);
}

deleteReplicationLink().catch(console.error);
