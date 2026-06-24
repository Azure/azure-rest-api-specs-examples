const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Kusto cluster.
 *
 * @summary update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClustersUpdate.json
 */
async function kustoClustersUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.update("kustorptest", "kustoCluster2", {}, { ifMatch: "*" });
  console.log(result);
}
