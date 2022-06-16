const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the network endpoints of all outbound dependencies of a Kusto cluster
 *
 * @summary Gets the network endpoints of all outbound dependencies of a Kusto cluster
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoOutboundNetworkDependenciesList.json
 */
async function getKustoClusterOutboundNetworkDependencies() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    clusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getKustoClusterOutboundNetworkDependencies().catch(console.error);
