const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a database's Advanced Threat Protection state.
 *
 * @summary creates or updates a database's Advanced Threat Protection state.
 * x-ms-original-file: 2025-01-01/DatabaseAdvancedThreatProtectionSettingsCreateMax.json
 */
async function updateADatabaseAdvancedThreatProtectionSettingsWithAllParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseAdvancedThreatProtectionSettings.createOrUpdate(
    "threatprotection-4799",
    "threatprotection-6440",
    "testdb",
    "Default",
    { state: "Enabled" },
  );
  console.log(result);
}
