const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the configurations of a cluster.
 *
 * @summary List all the configurations of a cluster.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ConfigurationListByCluster.json
 */
async function listConfigurationsOfTheCluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName =
    process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestResourceGroup";
  const clusterName = "testcluster";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurations.listByCluster(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
