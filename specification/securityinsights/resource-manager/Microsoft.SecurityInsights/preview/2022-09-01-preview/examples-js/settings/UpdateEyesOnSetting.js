const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates setting.
 *
 * @summary Updates setting.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/settings/UpdateEyesOnSetting.json
 */
async function updateEyesOnSettings() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const settingsName = "EyesOn";
  const settings = {
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    kind: "EyesOn",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.productSettings.update(
    resourceGroupName,
    workspaceName,
    settingsName,
    settings
  );
  console.log(result);
}
