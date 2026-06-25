const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to checks that the data connection parameters are valid.
 *
 * @summary checks that the data connection parameters are valid.
 * x-ms-original-file: 2025-02-14/KustoDataConnectionValidationAsync.json
 */
async function kustoDataConnectionValidation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.dataConnectionValidation(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase8",
    {
      dataConnectionName: "dataConnectionTest",
      properties: {
        kind: "EventHub",
        compression: "None",
        consumerGroup: "testConsumerGroup1",
        dataFormat: "MULTIJSON",
        eventHubResourceId:
          "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1",
        managedIdentityResourceId:
          "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
        mappingRuleName: "TestMapping",
        tableName: "TestTable",
      },
    },
  );
  console.log(result);
}
