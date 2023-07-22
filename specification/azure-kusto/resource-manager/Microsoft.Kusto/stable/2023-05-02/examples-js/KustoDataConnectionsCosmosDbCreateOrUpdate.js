const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a data connection.
 *
 * @summary Creates or updates a data connection.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDataConnectionsCosmosDbCreateOrUpdate.json
 */
async function kustoDataConnectionsCosmosDbCreateOrUpdate() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const databaseName = "KustoDatabase1";
  const dataConnectionName = "dataConnectionTest";
  const parameters = {
    cosmosDbAccountResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.DocumentDb/databaseAccounts/cosmosDbAccountTest1",
    cosmosDbContainer: "cosmosDbContainerTest",
    cosmosDbDatabase: "cosmosDbDatabaseTest",
    kind: "CosmosDb",
    location: "westus",
    managedIdentityResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/managedidentityTest1",
    mappingRuleName: "TestMapping",
    retrievalStartDate: new Date("2022-07-29T12:00:00.6554616Z"),
    tableName: "TestTable",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.dataConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    dataConnectionName,
    parameters
  );
  console.log(result);
}
