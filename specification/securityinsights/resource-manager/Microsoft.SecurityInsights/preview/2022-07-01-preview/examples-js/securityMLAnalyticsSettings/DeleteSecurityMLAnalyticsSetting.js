const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the Security ML Analytics Settings.
 *
 * @summary Delete the Security ML Analytics Settings.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/securityMLAnalyticsSettings/DeleteSecurityMLAnalyticsSetting.json
 */
async function deleteASecurityMlAnalyticsSettings() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const settingsResourceName = "f209187f-1d17-4431-94af-c141bf5f23db";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.securityMLAnalyticsSettings.delete(
    resourceGroupName,
    workspaceName,
    settingsResourceName
  );
  console.log(result);
}

deleteASecurityMlAnalyticsSettings().catch(console.error);
