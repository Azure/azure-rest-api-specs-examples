const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of recoverable databases
 *
 * @summary Gets a list of recoverable databases
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/RecoverableDatabaseList.json
 */
async function getListOfRestorableDroppedDatabases() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "recoverabledatabasetest-1234";
  const serverName = "recoverabledatabasetest-7177";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recoverableDatabases.listByServer(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getListOfRestorableDroppedDatabases().catch(console.error);
