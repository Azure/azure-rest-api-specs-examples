const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a server's Advanced Threat Protection state.
 *
 * @summary updates a server's Advanced Threat Protection state.
 * x-ms-original-file: 2025-06-01-preview/AdvancedThreatProtectionSettingsPatchEnabled.json
 */
async function enableAServerAdvancedThreatProtectionSettings() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.advancedThreatProtectionSettings.update(
    "threatprotection-4799",
    "threatprotection-6440",
    "Default",
    { state: "Enabled" },
  );
  console.log(result);
}
