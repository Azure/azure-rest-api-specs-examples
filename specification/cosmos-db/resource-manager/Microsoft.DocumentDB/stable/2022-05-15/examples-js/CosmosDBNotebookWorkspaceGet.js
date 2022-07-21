const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the notebook workspace for a Cosmos DB account.
 *
 * @summary Gets the notebook workspace for a Cosmos DB account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBNotebookWorkspaceGet.json
 */
async function cosmosDbNotebookWorkspaceGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const notebookWorkspaceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.notebookWorkspaces.get(
    resourceGroupName,
    accountName,
    notebookWorkspaceName
  );
  console.log(result);
}

cosmosDbNotebookWorkspaceGet().catch(console.error);
