const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves continuous backup information for a table.
 *
 * @summary Retrieves continuous backup information for a table.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBTableBackupInformation.json
 */
async function cosmosDbTableCollectionBackupInformation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const accountName = "ddb1";
  const tableName = "tableName1";
  const location = {
    location: "North Europe",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.tableResources.beginRetrieveContinuousBackupInformationAndWait(
    resourceGroupName,
    accountName,
    tableName,
    location
  );
  console.log(result);
}

cosmosDbTableCollectionBackupInformation().catch(console.error);
