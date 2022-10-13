const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the auth token for the notebook workspace
 *
 * @summary Regenerates the auth token for the notebook workspace
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBNotebookWorkspaceRegenerateAuthToken.json
 */
async function cosmosDbNotebookWorkspaceRegenerateAuthToken() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const notebookWorkspaceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.notebookWorkspaces.beginRegenerateAuthTokenAndWait(
    resourceGroupName,
    accountName,
    notebookWorkspaceName
  );
  console.log(result);
}

cosmosDbNotebookWorkspaceRegenerateAuthToken().catch(console.error);
