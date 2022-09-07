const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the FarmBeats instances for a subscription.
 *
 * @summary Lists the FarmBeats instances for a subscription.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_ListBySubscription.json
 */
async function farmBeatsModelsListBySubscription() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.farmBeatsModels.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

farmBeatsModelsListBySubscription().catch(console.error);
