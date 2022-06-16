const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of contact profiles by Subscription
 *
 * @summary Returns list of contact profiles by Subscription
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfilesBySubscriptionList.json
 */
async function listOfContactProfiles() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.contactProfiles.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfContactProfiles().catch(console.error);
