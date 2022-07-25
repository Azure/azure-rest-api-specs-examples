const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a workload classifier of Sql pool's workload group.
 *
 * @summary Get a workload classifier of Sql pool's workload group.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifier.json
 */
async function getAWorkloadClassifierForSqlAnalyticsPoolWorkloadGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const workloadGroupName = "wlm_workloadgroup";
  const workloadClassifierName = "wlm_classifier";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolWorkloadClassifier.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    workloadGroupName,
    workloadClassifierName
  );
  console.log(result);
}

getAWorkloadClassifierForSqlAnalyticsPoolWorkloadGroup().catch(console.error);
