const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the backups of this cluster that are available to restore.
 *
 * @summary List the backups of this cluster that are available to restore.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-05-15-preview/examples/CosmosDBManagedCassandraBackupsList.json
 */
async function cosmosDbManagedCassandraBackupsList() {
  const subscriptionId = "subid";
  const resourceGroupName = "cassandra-prod-rg";
  const clusterName = "cassandra-prod";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cassandraClusters.listBackups(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbManagedCassandraBackupsList().catch(console.error);
