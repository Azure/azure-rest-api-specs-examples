const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a data connection.
 *
 * @summary creates or updates a data connection.
 * x-ms-original-file: 2025-02-14/KustoDataConnectionsEventGridCreateOrUpdate.json
 */
async function kustoDataConnectionsEventGridCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.createOrUpdate(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase8",
    "dataConnectionTest",
    {
      kind: "EventGrid",
      location: "westus",
      blobStorageEventType: "Microsoft.Storage.BlobCreated",
      consumerGroup: "$Default",
      dataFormat: "MULTIJSON",
      databaseRouting: "Single",
      eventGridResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount/providers/Microsoft.EventGrid/eventSubscriptions/eventSubscriptionTest",
      eventHubResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest2",
      ignoreFirstRecord: false,
      managedIdentityResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
      mappingRuleName: "TestMapping",
      storageAccountResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/teststorageaccount",
      tableName: "TestTable",
    },
  );
  console.log(result);
}
