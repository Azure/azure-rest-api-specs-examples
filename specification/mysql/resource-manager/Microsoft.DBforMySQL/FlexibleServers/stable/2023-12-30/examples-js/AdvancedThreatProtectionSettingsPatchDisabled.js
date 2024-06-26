const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a server's Advanced Threat Protection state.
 *
 * @summary Updates a server's Advanced Threat Protection state.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/AdvancedThreatProtectionSettingsPatchDisabled.json
 */
async function disableAServerAdvancedThreatProtectionSettingsWithAllParameters() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "threatprotection-4799";
  const serverName = "threatprotection-6440";
  const advancedThreatProtectionName = "Default";
  const parameters = { state: "Disabled" };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.advancedThreatProtectionSettings.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    advancedThreatProtectionName,
    parameters,
  );
  console.log(result);
}
