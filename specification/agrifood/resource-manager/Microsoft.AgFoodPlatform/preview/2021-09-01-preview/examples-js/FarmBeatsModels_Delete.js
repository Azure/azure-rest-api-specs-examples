const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a FarmBeats resource.
 *
 * @summary Delete a FarmBeats resource.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_Delete.json
 */
async function farmBeatsModelsDelete() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmBeatsResourceName";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const result = await client.farmBeatsModels.delete(resourceGroupName, farmBeatsResourceName);
  console.log(result);
}

farmBeatsModelsDelete().catch(console.error);
