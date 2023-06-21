const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information of a configuration for coordinator and nodes.
 *
 * @summary Gets information of a configuration for coordinator and nodes.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationGet.json
 */
async function getConfigurationDetails() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName =
    process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestResourceGroup";
  const clusterName = "testcluster";
  const configurationName = "client_encoding";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.configurations.get(resourceGroupName, clusterName, configurationName);
  console.log(result);
}
