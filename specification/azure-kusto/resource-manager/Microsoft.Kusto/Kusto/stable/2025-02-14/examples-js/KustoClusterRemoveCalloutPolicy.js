const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to removes callout policy for engine services.
 *
 * @summary removes callout policy for engine services.
 * x-ms-original-file: 2025-02-14/KustoClusterRemoveCalloutPolicy.json
 */
async function kustoClusterDropCalloutPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  await client.clusters.removeCalloutPolicy("kustorptest", "kustoCluster", {
    calloutId: "*_kusto",
  });
}
