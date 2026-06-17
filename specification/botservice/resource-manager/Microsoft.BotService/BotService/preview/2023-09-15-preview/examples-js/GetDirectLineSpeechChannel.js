const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns a BotService Channel registration specified by the parameters.
 *
 * @summary returns a BotService Channel registration specified by the parameters.
 * x-ms-original-file: 2023-09-15-preview/GetDirectLineSpeechChannel.json
 */
async function getDirectLineSpeechChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.get(
    "OneResourceGroupName",
    "samplebotname",
    "DirectLineSpeechChannel",
  );
  console.log(result);
}
