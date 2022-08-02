const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the SIM groups in a subscription.
 *
 * @summary Gets all the SIM groups in a subscription.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupListBySubscription.json
 */
async function listSimGroupsInASubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.simGroups.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSimGroupsInASubscription().catch(console.error);
