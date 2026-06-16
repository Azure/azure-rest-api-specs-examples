const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Channel registration for a Bot Service
 *
 * @summary creates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/PutAlexaChannel.json
 */
async function createAlexaChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.create(
    "OneResourceGroupName",
    "samplebotname",
    "AlexaChannel",
    {
      location: "global",
      properties: {
        channelName: "AlexaChannel",
        properties: { alexaSkillId: "XAlexaSkillIdX", isEnabled: true },
      },
    },
  );
  console.log(result);
}
