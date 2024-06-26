const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SQL pool replication link by name.
 *
 * @summary Get SQL pool replication link by name.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolReplicationLinks_GetByName.json
 */
async function listsASqlAnalyticPoolReplicationLinks() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "sqlcrudtest-4799";
  const workspaceName = "sqlcrudtest-6440";
  const sqlPoolName = "testdb";
  const linkId = "5b301b68-03f6-4b26-b0f4-73ebb8634238";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolReplicationLinks.getByName(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    linkId
  );
  console.log(result);
}
