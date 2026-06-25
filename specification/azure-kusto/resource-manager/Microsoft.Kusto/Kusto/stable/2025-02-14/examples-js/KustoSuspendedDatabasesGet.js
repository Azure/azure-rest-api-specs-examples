const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns a database.
 *
 * @summary returns a database.
 * x-ms-original-file: 2025-02-14/KustoSuspendedDatabasesGet.json
 */
async function kustoSuspendedDatabasesGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.get("kustorptest", "kustoCluster", "KustoDatabase9");
  console.log(result);
}
