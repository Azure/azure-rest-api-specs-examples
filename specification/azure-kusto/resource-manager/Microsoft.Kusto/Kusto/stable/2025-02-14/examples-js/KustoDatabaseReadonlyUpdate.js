const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a database.
 *
 * @summary creates or updates a database.
 * x-ms-original-file: 2025-02-14/KustoDatabaseReadonlyUpdate.json
 */
async function kustoReadOnlyDatabaseUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.createOrUpdate(
    "kustorptest",
    "kustoCluster",
    "kustoReadOnlyDatabase",
    { kind: "ReadOnlyFollowing", location: "westus", hotCachePeriod: "P1D" },
  );
  console.log(result);
}
