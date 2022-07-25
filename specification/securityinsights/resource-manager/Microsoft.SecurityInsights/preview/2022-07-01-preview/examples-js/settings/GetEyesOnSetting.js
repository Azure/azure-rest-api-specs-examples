const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a setting.
 *
 * @summary Gets a setting.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/settings/GetEyesOnSetting.json
 */
async function getEyesOnSettings() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const settingsName = "EyesOn";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.productSettings.get(resourceGroupName, workspaceName, settingsName);
  console.log(result);
}

getEyesOnSettings().catch(console.error);
