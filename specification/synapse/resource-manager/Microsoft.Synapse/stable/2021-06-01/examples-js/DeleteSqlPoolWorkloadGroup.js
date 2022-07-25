const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove Sql pool's workload group.
 *
 * @summary Remove Sql pool's workload group.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroup.json
 */
async function deleteAWorkloadGroupOfASqlAnalyticsPool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const workloadGroupName = "wlm_workloadgroup";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolWorkloadGroup.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    workloadGroupName
  );
  console.log(result);
}

deleteAWorkloadGroupOfASqlAnalyticsPool().catch(console.error);
