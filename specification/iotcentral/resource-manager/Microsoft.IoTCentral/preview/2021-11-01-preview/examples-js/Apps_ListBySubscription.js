const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all IoT Central Applications in a subscription.
 *
 * @summary Get all IoT Central Applications in a subscription.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_ListBySubscription.json
 */
async function appsListBySubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.apps.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

appsListBySubscription().catch(console.error);
