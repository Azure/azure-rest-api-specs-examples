const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a service.
 *
 * @summary Creates a service.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBSqlDedicatedGatewayServiceCreate.json
 */
async function sqlDedicatedGatewayServiceCreate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const serviceName = "SqlDedicatedGateway";
  const createUpdateParameters = {
    instanceCount: 1,
    instanceSize: "Cosmos.D4s",
    serviceType: "SqlDedicatedGateway",
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
