const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List schedules in a subscription.
 *
 * @summary List schedules in a subscription.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_ListBySubscription.json
 */
async function globalSchedulesListBySubscription() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.globalSchedules.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

globalSchedulesListBySubscription().catch(console.error);
