const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of database scripts for given database.
 *
 * @summary Returns the list of database scripts for given database.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoScriptsListByDatabase.json
 */
async function kustoScriptsList() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "Kustodatabase8";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scripts.listByDatabase(
    resourceGroupName,
    clusterName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoScriptsList().catch(console.error);
