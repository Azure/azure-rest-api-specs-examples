const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a SQL pool
 *
 * @summary Create a SQL pool
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
 */
async function createASqlAnalyticsPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const sqlPoolName = "ExampleSqlPool";
  const sqlPoolInfo = {
    collation: "",
    createMode: "",
    location: "Southeast Asia",
    maxSizeBytes: 0,
    recoverableDatabaseId: "",
    sku: { name: "", tier: "" },
    sourceDatabaseId: "",
    storageAccountType: "LRS",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPools.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    sqlPoolInfo
  );
  console.log(result);
}
