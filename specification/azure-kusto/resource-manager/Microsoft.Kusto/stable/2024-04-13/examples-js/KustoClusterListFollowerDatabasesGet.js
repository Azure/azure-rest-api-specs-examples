const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns a list of databases that are owned by this cluster and were followed by another cluster.
 *
 * @summary Returns a list of databases that are owned by this cluster and were followed by another cluster.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClusterListFollowerDatabasesGet.json
 */
async function kustoClusterListFollowerDatabasesGet() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listFollowerDatabasesGet(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
