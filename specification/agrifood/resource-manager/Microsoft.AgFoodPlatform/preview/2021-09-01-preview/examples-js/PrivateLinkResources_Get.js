const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Private link resource object.
 *
 * @summary Get Private link resource object.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateLinkResources_Get.json
 */
async function privateLinkResourcesGet() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmbeatsResourceName";
  const subResourceName = "farmbeats";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(
    resourceGroupName,
    farmBeatsResourceName,
    subResourceName
  );
  console.log(result);
}

privateLinkResourcesGet().catch(console.error);
