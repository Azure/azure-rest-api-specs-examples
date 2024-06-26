const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a service.
 *
 * @summary Creates a service.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBMaterializedViewsBuilderServiceCreate.json
 */
async function materializedViewsBuilderServiceCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const serviceName = "MaterializedViewsBuilder";
  const createUpdateParameters = {
    instanceCount: 1,
    instanceSize: "Cosmos.D4s",
    serviceType: "MaterializedViewsBuilder",
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

materializedViewsBuilderServiceCreate().catch(console.error);
