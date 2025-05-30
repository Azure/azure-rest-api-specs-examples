const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an attached database configuration.
 *
 * @summary Creates or updates an attached database configuration.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
 */
async function attachedDatabaseConfigurationsCreateOrUpdate() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster2";
  const attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
  const parameters = {
    clusterResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2",
    databaseName: "kustodatabase",
    databaseNameOverride: "overridekustodatabase",
    defaultPrincipalsModificationKind: "Union",
    location: "westus",
    tableLevelSharingProperties: {
      externalTablesToExclude: ["ExternalTable2"],
      externalTablesToInclude: ["ExternalTable1"],
      materializedViewsToExclude: ["MaterializedViewTable2"],
      materializedViewsToInclude: ["MaterializedViewTable1"],
      tablesToExclude: ["Table2"],
      tablesToInclude: ["Table1"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.attachedDatabaseConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    attachedDatabaseConfigurationName,
    parameters,
  );
  console.log(result);
}
