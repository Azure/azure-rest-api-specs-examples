const { HealthbotClient } = require("@azure/arm-healthbot");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all the resources of a particular type belonging to a subscription.
 *
 * @summary Returns all the resources of a particular type belonging to a subscription.
 * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ListBotsBySubscription.json
 */
async function listBotsBySubscription() {
  const subscriptionId = "subscription-id";
  const credential = new DefaultAzureCredential();
  const client = new HealthbotClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBotsBySubscription().catch(console.error);
