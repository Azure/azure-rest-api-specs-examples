const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the config reference and status of an app
 *
 * @summary Description for Gets the config reference and status of an app
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSetting.json
 */
async function getAzureKeyVaultAppSettingReference() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "testc6282";
  const appSettingKey = "setting";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.getAppSettingKeyVaultReference(
    resourceGroupName,
    name,
    appSettingKey
  );
  console.log(result);
}

getAzureKeyVaultAppSettingReference().catch(console.error);
