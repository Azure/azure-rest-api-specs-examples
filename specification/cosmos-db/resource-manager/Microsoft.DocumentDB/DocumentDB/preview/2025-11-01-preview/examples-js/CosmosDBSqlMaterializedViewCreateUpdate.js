const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an Azure Cosmos DB SQL container
 *
 * @summary create or update an Azure Cosmos DB SQL container
 * x-ms-original-file: 2025-11-01-preview/CosmosDBSqlMaterializedViewCreateUpdate.json
 */
async function cosmosDBSqlMaterializedViewCreateUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.createUpdateSqlContainer(
    "rg1",
    "ddb1",
    "databaseName",
    "mvContainerName",
    {
      location: "West US",
      tags: {},
      resource: {
        id: "mvContainerName",
        indexingPolicy: {
          indexingMode: "consistent",
          automatic: true,
          includedPaths: [
            {
              path: "/*",
              indexes: [
                { kind: "Range", dataType: "String", precision: -1 },
                { kind: "Range", dataType: "Number", precision: -1 },
              ],
            },
          ],
          excludedPaths: [],
        },
        partitionKey: { paths: ["/mvpk"], kind: "Hash" },
        materializedViewDefinition: {
          sourceCollectionId: "sourceContainerName",
          definition: "select * from ROOT",
          throughputBucketForBuild: 1,
        },
      },
      options: {},
    },
  );
  console.log(result);
}
