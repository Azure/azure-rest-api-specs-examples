const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the database with the given name.
 *
 * @summary Deletes the database with the given name.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabasesDelete.json
 */
async function kustoDatabasesDelete() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase8";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    databaseName
  );
  console.log(result);
}

kustoDatabasesDelete().catch(console.error);
