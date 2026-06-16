const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Channel registration for a Bot Service
 *
 * @summary updates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/UpdateAlexaChannel.json
 */
async function updateAlexaChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.update(
    "OneResourceGroupName",
    "samplebotname",
    "AlexaChannel",
    {
      properties: {
        channelName: "AlexaChannel",
        properties: { alexaSkillId: "XAlexaSkillIdX", isEnabled: true },
      },
      location: "global",
    },
  );
  console.log(result);
}
