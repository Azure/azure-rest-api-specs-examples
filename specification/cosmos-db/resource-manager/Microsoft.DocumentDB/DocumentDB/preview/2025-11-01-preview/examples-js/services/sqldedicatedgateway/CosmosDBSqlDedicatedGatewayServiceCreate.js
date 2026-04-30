const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a service.
 *
 * @summary creates a service.
 * x-ms-original-file: 2025-11-01-preview/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceCreate.json
 */
async function sqlDedicatedGatewayServiceCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.service.create("rg1", "ddb1", "SqlDedicatedGateway", {
    properties: {
      dedicatedGatewayType: "IntegratedCache",
      instanceCount: 1,
      instanceSize: "Cosmos.D4s",
      serviceType: "SqlDedicatedGateway",
    },
  });
  console.log(result);
}
