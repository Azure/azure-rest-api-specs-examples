const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the attached database configuration with the given name.
 *
 * @summary Deletes the attached database configuration with the given name.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsDelete.json
 */
async function attachedDatabaseConfigurationsDelete() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.attachedDatabaseConfigurations.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    attachedDatabaseConfigurationName
  );
  console.log(result);
}
