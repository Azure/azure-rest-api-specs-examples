const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Applies the configuration settings from the target slot onto the current slot.
 *
 * @summary Description for Applies the configuration settings from the target slot onto the current slot.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ApplySlotConfig.json
 */
async function applyWebAppSlotConfig() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "sitef6141";
  const slotSwapEntity = {
    preserveVnet: true,
    targetSlot: "staging",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.applySlotConfigToProduction(
    resourceGroupName,
    name,
    slotSwapEntity
  );
  console.log(result);
}
