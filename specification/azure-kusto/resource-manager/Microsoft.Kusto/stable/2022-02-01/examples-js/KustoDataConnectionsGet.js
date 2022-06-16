const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a data connection.
 *
 * @summary Returns a data connection.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsGet.json
 */
async function kustoDataConnectionsGet() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase8";
  const dataConnectionName = "dataConnectionTest";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.get(
    resourceGroupName,
    clusterName,
    databaseName,
    dataConnectionName
  );
  console.log(result);
}

kustoDataConnectionsGet().catch(console.error);
