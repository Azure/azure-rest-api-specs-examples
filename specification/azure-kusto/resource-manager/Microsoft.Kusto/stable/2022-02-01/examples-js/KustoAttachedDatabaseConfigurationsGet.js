const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns an attached database configuration.
 *
 * @summary Returns an attached database configuration.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsGet.json
 */
async function attachedDatabaseConfigurationsGet() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const clusterName = "kustoCluster2";
  const attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.attachedDatabaseConfigurations.get(
    resourceGroupName,
    clusterName,
    attachedDatabaseConfigurationName
  );
  console.log(result);
}

attachedDatabaseConfigurationsGet().catch(console.error);
