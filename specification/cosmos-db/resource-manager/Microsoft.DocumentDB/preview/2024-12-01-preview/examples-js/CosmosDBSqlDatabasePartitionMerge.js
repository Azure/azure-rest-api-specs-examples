const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Merges the partitions of a SQL database
 *
 * @summary Merges the partitions of a SQL database
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlDatabasePartitionMerge.json
 */
async function cosmosDbSqlDatabasePartitionMerge() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const mergeParameters = { isDryRun: false };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginSqlDatabasePartitionMergeAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    mergeParameters,
  );
  console.log(result);
}
