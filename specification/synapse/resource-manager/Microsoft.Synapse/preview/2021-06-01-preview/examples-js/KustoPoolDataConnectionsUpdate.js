const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a data connection.
 *
 * @summary Updates a data connection.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsUpdate.json
 */
async function kustoPoolDataConnectionsUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const dataConnectionName = "DataConnections8";
  const parameters = {
    consumerGroup: "testConsumerGroup1",
    eventHubResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1",
    kind: "EventHub",
    location: "westus",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDataConnections.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    dataConnectionName,
    parameters
  );
  console.log(result);
}

kustoPoolDataConnectionsUpdate().catch(console.error);
