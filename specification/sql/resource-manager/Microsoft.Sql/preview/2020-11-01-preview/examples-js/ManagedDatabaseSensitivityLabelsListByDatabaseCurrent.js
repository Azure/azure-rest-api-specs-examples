const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the sensitivity labels of a given database
 *
 * @summary Gets the sensitivity labels of a given database
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsListByDatabaseCurrent.json
 */
async function getsTheCurrentSensitivityLabelsOfAGivenDatabaseInAManagedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseSensitivityLabels.listCurrentByDatabase(
    resourceGroupName,
    managedInstanceName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheCurrentSensitivityLabelsOfAGivenDatabaseInAManagedDatabase().catch(console.error);
