const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a database.
 *
 * @summary Updates a database.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesUpdate.json
 */
async function kustoPoolDatabasesUpdate() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const parameters = {
    kind: "ReadWrite",
    softDeletePeriod: "P1D",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDatabases.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    parameters
  );
  console.log(result);
}
