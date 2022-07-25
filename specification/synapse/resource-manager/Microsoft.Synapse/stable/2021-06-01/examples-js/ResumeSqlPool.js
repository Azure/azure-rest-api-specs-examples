const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resume a SQL pool
 *
 * @summary Resume a SQL pool
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ResumeSqlPool.json
 */
async function resumeASqlAnalyticsPool() {
  const subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-9187";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPools.beginResumeAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName
  );
  console.log(result);
}

resumeASqlAnalyticsPool().catch(console.error);
