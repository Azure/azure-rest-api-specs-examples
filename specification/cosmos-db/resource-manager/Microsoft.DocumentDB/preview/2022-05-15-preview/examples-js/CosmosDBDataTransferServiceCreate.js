const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a service.
 *
 * @summary Creates a service.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBDataTransferServiceCreate.json
 */
async function dataTransferServiceCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const serviceName = "DataTransfer";
  const createUpdateParameters = {
    instanceCount: 1,
    instanceSize: "Cosmos.D4s",
    serviceType: "DataTransfer",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.service.beginCreateAndWait(
    resourceGroupName,
    accountName,
    serviceName,
    createUpdateParameters
  );
  console.log(result);
}

dataTransferServiceCreate().catch(console.error);
