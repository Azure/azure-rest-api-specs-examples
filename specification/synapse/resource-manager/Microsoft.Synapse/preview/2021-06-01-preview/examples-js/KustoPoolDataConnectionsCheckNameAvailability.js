const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the data connection name is valid and is not already in use.
 *
 * @summary Checks that the data connection name is valid and is not already in use.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCheckNameAvailability.json
 */
async function kustoPoolDataConnectionsCheckNameAvailability() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "Kustodatabase8";
  const dataConnectionName = {
    name: "DataConnections8",
    type: "Microsoft.Synapse/workspaces/kustoPools/databases/dataConnections",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDataConnections.checkNameAvailability(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    dataConnectionName
  );
  console.log(result);
}
