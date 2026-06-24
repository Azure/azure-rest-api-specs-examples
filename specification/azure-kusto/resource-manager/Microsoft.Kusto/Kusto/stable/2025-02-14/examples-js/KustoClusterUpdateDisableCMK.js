const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Kusto cluster.
 *
 * @summary update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClusterUpdateDisableCMK.json
 */
async function kustoClusterUpdateDisableCMK() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.update(
    "kustoRgTest",
    "kustoClusterCMK",
    { location: "westus", keyVaultProperties: {} },
    { ifMatch: "*" },
  );
  console.log(result);
}
