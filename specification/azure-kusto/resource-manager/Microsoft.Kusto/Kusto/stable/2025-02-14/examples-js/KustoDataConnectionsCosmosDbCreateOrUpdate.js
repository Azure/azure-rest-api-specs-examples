const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a data connection.
 *
 * @summary creates or updates a data connection.
 * x-ms-original-file: 2025-02-14/KustoDataConnectionsCosmosDbCreateOrUpdate.json
 */
async function kustoDataConnectionsCosmosDbCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.createOrUpdate(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase1",
    "dataConnectionTest",
    {
      kind: "CosmosDb",
      location: "westus",
      cosmosDbAccountResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1",
      cosmosDbContainer: "cosmosDbContainerTest",
      cosmosDbDatabase: "cosmosDbDatabaseTest",
      managedIdentityResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
      mappingRuleName: "TestMapping",
      retrievalStartDate: new Date("2022-07-29T12:00:00.6554616Z"),
      tableName: "TestTable",
    },
  );
  console.log(result);
}
