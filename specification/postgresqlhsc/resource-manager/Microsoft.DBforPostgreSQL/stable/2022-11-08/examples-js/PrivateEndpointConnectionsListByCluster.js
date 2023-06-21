const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets list of private endpoint connections on a cluster.
 *
 * @summary Gets list of private endpoint connections on a cluster.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateEndpointConnectionsListByCluster.json
 */
async function getsListOfPrivateEndpointConnectionsOnACluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName =
    process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestResourceGroup";
  const clusterName = "testcluster";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByCluster(
    resourceGroupName,
    clusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
