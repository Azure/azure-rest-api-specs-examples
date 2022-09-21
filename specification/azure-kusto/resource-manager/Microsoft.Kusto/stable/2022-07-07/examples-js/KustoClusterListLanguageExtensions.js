const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of language extensions that can run within KQL queries.
 *
 * @summary Returns a list of language extensions that can run within KQL queries.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClusterListLanguageExtensions.json
 */
async function kustoClusterListLanguageExtensions() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listLanguageExtensions(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoClusterListLanguageExtensions().catch(console.error);
