const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a database.
 *
 * @summary Creates or updates a database.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesCreateOrUpdate.json
 */
async function kustoPoolDatabasesCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const parameters = {
    kind: "ReadWrite",
    location: "westus",
    softDeletePeriod: "P1D",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDatabases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    parameters
  );
  console.log(result);
}

kustoPoolDatabasesCreateOrUpdate().catch(console.error);
