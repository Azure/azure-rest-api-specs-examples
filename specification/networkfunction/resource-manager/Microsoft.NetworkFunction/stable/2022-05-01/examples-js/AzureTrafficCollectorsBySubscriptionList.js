const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of Azure Traffic Collectors in a subscription
 *
 * @summary Return list of Azure Traffic Collectors in a subscription
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorsBySubscriptionList.json
 */
async function listOfTrafficCollectorsBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureTrafficCollectorsBySubscription.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfTrafficCollectorsBySubscription().catch(console.error);
