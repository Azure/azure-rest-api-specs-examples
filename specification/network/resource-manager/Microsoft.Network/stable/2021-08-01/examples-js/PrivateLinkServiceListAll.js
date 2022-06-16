const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private link service in a subscription.
 *
 * @summary Gets all private link service in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateLinkServiceListAll.json
 */
async function listAllPrivateListService() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllPrivateListService().catch(console.error);
