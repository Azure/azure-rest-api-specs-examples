Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-kusto_7.1.1/sdk/kusto/arm-kusto/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an attached database configuration.
 *
 * @summary Creates or updates an attached database configuration.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
 */
async function attachedDatabaseConfigurationsCreateOrUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster2";
  const attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
  const parameters = {
    clusterResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2",
    databaseName: "kustodatabase",
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
    parameters
  );
  console.log(result);
}

attachedDatabaseConfigurationsCreateOrUpdate().catch(console.error);
```
