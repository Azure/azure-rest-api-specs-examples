const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Kusto cluster.
 *
 * @summary Deletes a Kusto cluster.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClustersDelete.json
 */
async function kustoClustersDelete() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster2";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginDeleteAndWait(resourceGroupName, clusterName);
  console.log(result);
}

kustoClustersDelete().catch(console.error);
