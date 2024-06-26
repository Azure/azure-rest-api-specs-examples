const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get installed extensions details.
 *
 * @summary Get installed extensions details.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_ListByFarmBeats.json
 */
async function extensionsListByFarmBeats() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmbeatsResourceName";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extensions.listByFarmBeats(
    resourceGroupName,
    farmBeatsResourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

extensionsListByFarmBeats().catch(console.error);
