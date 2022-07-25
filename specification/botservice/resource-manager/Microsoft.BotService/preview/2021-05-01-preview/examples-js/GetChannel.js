const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a BotService Channel registration specified by the parameters.
 *
 * @summary Returns a BotService Channel registration specified by the parameters.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetChannel.json
 */
async function getBot() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "EmailChannel";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.get(resourceGroupName, resourceName, channelName);
  console.log(result);
}

getBot().catch(console.error);
