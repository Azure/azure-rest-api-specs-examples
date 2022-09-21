const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a database.
 *
 * @summary Creates or updates a database.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDatabasesCreateOrUpdate.json
 */
async function kustoReadWriteDatabaseCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase8";
  const callerRole = "Admin";
  const parameters = {
    kind: "ReadWrite",
    location: "westus",
    softDeletePeriod: "P1D",
  };
  const options = { callerRole };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters,
    options
  );
  console.log(result);
}

kustoReadWriteDatabaseCreateOrUpdate().catch(console.error);
