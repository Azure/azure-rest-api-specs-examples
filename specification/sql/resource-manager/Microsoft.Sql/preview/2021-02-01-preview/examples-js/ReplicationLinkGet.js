const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a replication link.
 *
 * @summary Gets a replication link.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ReplicationLinkGet.json
 */
async function getsTheReplicationLink() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "sourcesvr";
  const databaseName = "gamma-db";
  const linkId = "4891ca10-ebd0-47d7-9182-c722651780fb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.replicationLinks.get(
    resourceGroupName,
    serverName,
    databaseName,
    linkId
  );
  console.log(result);
}

getsTheReplicationLink().catch(console.error);
