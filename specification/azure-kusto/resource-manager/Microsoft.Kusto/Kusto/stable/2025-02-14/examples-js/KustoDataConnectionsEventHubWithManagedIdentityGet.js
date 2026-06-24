const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns a data connection.
 *
 * @summary returns a data connection.
 * x-ms-original-file: 2025-02-14/KustoDataConnectionsEventHubWithManagedIdentityGet.json
 */
async function kustoDataConnectionsEventHubWithManagedIdentityGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.get(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase8",
    "dataConnectionTest",
  );
  console.log(result);
}
