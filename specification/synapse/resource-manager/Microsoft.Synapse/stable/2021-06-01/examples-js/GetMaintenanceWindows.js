const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a SQL pool's Maintenance Windows.
 *
 * @summary Get a SQL pool's Maintenance Windows.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetMaintenanceWindows.json
 */
async function getsMaintenanceWindowSettingsForASelectedSqlAnalyticsPool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "samplerg";
  const workspaceName = "testworkspace";
  const sqlPoolName = "testsp";
  const maintenanceWindowName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolMaintenanceWindows.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    maintenanceWindowName
  );
  console.log(result);
}

getsMaintenanceWindowSettingsForASelectedSqlAnalyticsPool().catch(console.error);
