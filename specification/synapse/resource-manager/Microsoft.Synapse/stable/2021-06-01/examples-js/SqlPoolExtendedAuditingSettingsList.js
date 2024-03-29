const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists extended auditing settings of a Sql pool.
 *
 * @summary Lists extended auditing settings of a Sql pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolExtendedAuditingSettingsList.json
 */
async function listExtendedAuditingSettingsOfADatabase() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "blobauditingtest-6852";
  const workspaceName = "blobauditingtest-2080";
  const sqlPoolName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extendedSqlPoolBlobAuditingPolicies.listBySqlPool(
    resourceGroupName,
    workspaceName,
    sqlPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
