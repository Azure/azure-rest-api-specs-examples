const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Apply a partial update to a SQL pool
 *
 * @summary Apply a partial update to a SQL pool
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
 */
async function updateASqlAnalyticsPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const sqlPoolName = "ExampleSqlPool";
  const sqlPoolInfo = {
    collation: "",
    location: "West US 2",
    maxSizeBytes: 0,
    restorePointInTime: new Date("1970-01-01T00:00:00.000Z"),
    sku: { name: "", tier: "" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPools.update(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    sqlPoolInfo
  );
  console.log(result);
}
