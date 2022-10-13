const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the usages (most recent storage data) for the given collection, split by partition.
 *
 * @summary Retrieves the usages (most recent storage data) for the given collection, split by partition.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBCollectionPartitionGetUsages.json
 */
async function cosmosDbCollectionGetUsages() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseRid = "databaseRid";
  const collectionRid = "collectionRid";
  const filter = "$filter=name.value eq 'Partition Storage'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.collectionPartition.listUsages(
    resourceGroupName,
    accountName,
    databaseRid,
    collectionRid,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbCollectionGetUsages().catch(console.error);
