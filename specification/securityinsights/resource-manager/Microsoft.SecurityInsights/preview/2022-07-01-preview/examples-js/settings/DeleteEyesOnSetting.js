const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete setting of the product.
 *
 * @summary Delete setting of the product.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/settings/DeleteEyesOnSetting.json
 */
async function deleteEyesOnSettings() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const settingsName = "EyesOn";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.productSettings.delete(
    resourceGroupName,
    workspaceName,
    settingsName
  );
  console.log(result);
}

deleteEyesOnSettings().catch(console.error);
