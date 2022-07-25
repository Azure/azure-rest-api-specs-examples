const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB Graph.
 *
 * @summary Create or update an Azure Cosmos DB Graph.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBGraphResourceCreateUpdate.json
 */
async function cosmosDbGraphCreateUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const graphName = "graphName";
  const createUpdateGraphParameters = {
    location: "West US",
    options: {},
    resource: { id: "graphName" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.graphResources.beginCreateUpdateGraphAndWait(
    resourceGroupName,
    accountName,
    graphName,
    createUpdateGraphParameters
  );
  console.log(result);
}

cosmosDbGraphCreateUpdate().catch(console.error);
