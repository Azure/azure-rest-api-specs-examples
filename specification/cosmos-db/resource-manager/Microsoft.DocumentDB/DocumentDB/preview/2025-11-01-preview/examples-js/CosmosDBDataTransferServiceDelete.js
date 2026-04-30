const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes service with the given serviceName.
 *
 * @summary deletes service with the given serviceName.
 * x-ms-original-file: 2025-11-01-preview/CosmosDBDataTransferServiceDelete.json
 */
async function dataTransferServiceDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  await client.service.delete("rg1", "ddb1", "DataTransfer");
}
