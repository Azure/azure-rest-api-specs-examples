const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a data connection.
 *
 * @summary updates a data connection.
 * x-ms-original-file: 2025-02-14/KustoDataConnectionsEventHubWithManagedIdentityUpdate.json
 */
async function kustoDataConnectionsEventHubWithManagedIdentityUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.update(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase8",
    "dataConnectionTest",
    {
      location: "westus",
      kind: "EventHubWithManagedIdentity",
      eventHubResourceIdForManagedIdentity:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1",
      managedIdentityResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
      consumerGroup: "testConsumerGroup1",
    },
  );
  console.log(result);
}
