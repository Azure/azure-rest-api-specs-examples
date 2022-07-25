const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Merges the partitions of a SQL Container
 *
 * @summary Merges the partitions of a SQL Container
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBSqlContainerPartitionMerge.json
 */
async function cosmosDbSqlContainerPartitionMerge() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const mergeParameters = { isDryRun: false };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginListSqlContainerPartitionMergeAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
    mergeParameters
  );
  console.log(result);
}

cosmosDbSqlContainerPartitionMerge().catch(console.error);
