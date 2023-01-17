const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all the resources of a particular type belonging to a subscription.
 *
 * @summary Returns all the resources of a particular type belonging to a subscription.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListBotsBySubscription.json
 */
async function listBotsBySubscription() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
