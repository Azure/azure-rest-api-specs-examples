const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Data Transfer Job.
 *
 * @summary Get a Data Transfer Job.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/data-transfer-service/CosmosDBDataTransferJobGet.json
 */
async function cosmosDbDataTransferJobGet() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const jobName = "j1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.dataTransferJobs.get(resourceGroupName, accountName, jobName);
  console.log(result);
}
