const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the sims in a subscription.
 *
 * @summary Gets all the sims in a subscription.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimListBySubscription.json
 */
async function listSimsInASubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sims.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSimsInASubscription().catch(console.error);
