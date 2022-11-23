const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Microsoft.DevCenter SKUs available in a subscription
 *
 * @summary Lists the Microsoft.DevCenter SKUs available in a subscription
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/Skus_ListBySubscription.json
 */
async function skusListBySubscription() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skus.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

skusListBySubscription().catch(console.error);
