const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts the notebook workspace
 *
 * @summary Starts the notebook workspace
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBNotebookWorkspaceStart.json
 */
async function cosmosDbNotebookWorkspaceStart() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const notebookWorkspaceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.notebookWorkspaces.beginStartAndWait(
    resourceGroupName,
    accountName,
    notebookWorkspaceName,
  );
  console.log(result);
}
