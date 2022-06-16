const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the long term retention backups for managed databases in a given location.
 *
 * @summary Lists the long term retention backups for managed databases in a given location.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceLongTermRetentionBackupListByLocation.json
 */
async function getAllLongTermRetentionBackupsUnderTheLocation() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "japaneast";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.longTermRetentionManagedInstanceBackups.listByLocation(
    locationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllLongTermRetentionBackupsUnderTheLocation().catch(console.error);
