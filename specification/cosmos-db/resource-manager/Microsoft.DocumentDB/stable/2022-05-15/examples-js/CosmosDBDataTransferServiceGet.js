const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of service.
 *
 * @summary Gets the status of service.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDataTransferServiceGet.json
 */
async function dataTransferServiceGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const serviceName = "DataTransfer";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.service.get(resourceGroupName, accountName, serviceName);
  console.log(result);
}

dataTransferServiceGet().catch(console.error);
