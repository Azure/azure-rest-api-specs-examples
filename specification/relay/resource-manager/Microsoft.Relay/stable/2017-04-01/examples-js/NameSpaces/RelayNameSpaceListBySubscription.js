const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available namespaces within the subscription regardless of the resourceGroups.
 *
 * @summary Lists all the available namespaces within the subscription regardless of the resourceGroups.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceListBySubscription.json
 */
async function relayNameSpaceListBySubscription() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
