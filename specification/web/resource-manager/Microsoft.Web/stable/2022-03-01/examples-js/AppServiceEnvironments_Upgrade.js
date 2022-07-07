const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Initiate an upgrade of an App Service Environment if one is available.
 *
 * @summary Description for Initiate an upgrade of an App Service Environment if one is available.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_Upgrade.json
 */
async function initiateAnUpgradeOnAnAppServiceEnvironment() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "SampleHostingEnvironment";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceEnvironments.beginUpgradeAndWait(resourceGroupName, name);
  console.log(result);
}

initiateAnUpgradeOnAnAppServiceEnvironment().catch(console.error);
