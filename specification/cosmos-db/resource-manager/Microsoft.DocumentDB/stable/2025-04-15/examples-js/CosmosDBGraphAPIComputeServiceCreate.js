const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a service.
 *
 * @summary Creates a service.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBGraphAPIComputeServiceCreate.json
 */
async function graphApiComputeServiceCreate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const serviceName = "GraphAPICompute";
  const createUpdateParameters = {
    properties: {
      instanceCount: 1,
      instanceSize: "Cosmos.D4s",
      serviceType: "GraphAPICompute",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.service.beginCreateAndWait(
    resourceGroupName,
    accountName,
    serviceName,
    createUpdateParameters,
  );
  console.log(result);
}
