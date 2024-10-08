const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fails over from the current primary server to this server allowing data loss.
 *
 * @summary Fails over from the current primary server to this server allowing data loss.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ReplicationLinkFailoverAllowDataLoss.json
 */
async function forcedFailoverOfAReplicationLink() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default";
  const serverName = "sourcesvr";
  const databaseName = "gamma-db";
  const linkId = "4891ca10-ebd0-47d7-9182-c722651780fb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.replicationLinks.beginFailoverAllowDataLossAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    linkId,
  );
  console.log(result);
}
