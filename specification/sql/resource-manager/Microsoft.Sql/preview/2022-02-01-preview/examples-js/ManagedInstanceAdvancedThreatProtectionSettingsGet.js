const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a managed instance's Advanced Threat Protection state.
 *
 * @summary Get a managed instance's Advanced Threat Protection state.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedInstanceAdvancedThreatProtectionSettingsGet.json
 */
async function getAManagedInstanceAdvancedThreatProtectionSettings() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "threatprotection-4799";
  const managedInstanceName = "threatprotection-6440";
  const advancedThreatProtectionName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceAdvancedThreatProtectionSettings.get(
    resourceGroupName,
    managedInstanceName,
    advancedThreatProtectionName,
  );
  console.log(result);
}
