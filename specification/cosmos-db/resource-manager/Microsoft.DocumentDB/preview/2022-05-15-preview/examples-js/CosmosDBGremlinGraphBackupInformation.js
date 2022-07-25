const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves continuous backup information for a gremlin graph.
 *
 * @summary Retrieves continuous backup information for a gremlin graph.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBGremlinGraphBackupInformation.json
 */
async function cosmosDbGremlinGraphBackupInformation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const graphName = "graphName";
  const location = {
    location: "North Europe",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.gremlinResources.beginRetrieveContinuousBackupInformationAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    graphName,
    location
  );
  console.log(result);
}

cosmosDbGremlinGraphBackupInformation().catch(console.error);
