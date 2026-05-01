const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the status of service.
 *
 * @summary gets the status of service.
 * x-ms-original-file: 2025-11-01-preview/CosmosDBMaterializedViewsBuilderServiceGet.json
 */
async function materializedViewsBuilderServiceGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.service.get("rg1", "ddb1", "MaterializedViewsBuilder");
  console.log(result);
}
