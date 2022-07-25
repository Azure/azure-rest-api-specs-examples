const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Channel registration from a Bot Service
 *
 * @summary Deletes a Channel registration from a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/DeleteDirectLineSpeechChannel.json
 */
async function deleteDirectLineSpeechBot() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "DirectLineSpeechChannel";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.delete(resourceGroupName, resourceName, channelName);
  console.log(result);
}

deleteDirectLineSpeechBot().catch(console.error);
