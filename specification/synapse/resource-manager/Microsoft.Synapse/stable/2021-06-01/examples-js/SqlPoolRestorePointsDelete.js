const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a restore point.
 *
 * @summary Deletes a restore point.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsDelete.json
 */
async function deletesARestorePoint() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const workspaceName = "testws";
  const sqlPoolName = "testpool";
  const restorePointName = "131546477590000000";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolRestorePoints.delete(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    restorePointName
  );
  console.log(result);
}
