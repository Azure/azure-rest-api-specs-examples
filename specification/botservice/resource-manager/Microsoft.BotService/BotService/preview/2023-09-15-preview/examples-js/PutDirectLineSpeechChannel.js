const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Channel registration for a Bot Service
 *
 * @summary creates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/PutDirectLineSpeechChannel.json
 */
async function createDirectLineSpeechChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.create(
    "OneResourceGroupName",
    "samplebotname",
    "DirectLineSpeechChannel",
    {
      location: "global",
      properties: {
        channelName: "DirectLineSpeechChannel",
        properties: {
          cognitiveServiceRegion: "XcognitiveServiceRegionX",
          cognitiveServiceSubscriptionKey: "XcognitiveServiceSubscriptionKeyX",
          isEnabled: true,
        },
      },
    },
  );
  console.log(result);
}
