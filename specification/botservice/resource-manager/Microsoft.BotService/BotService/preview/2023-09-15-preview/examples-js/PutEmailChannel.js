const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Channel registration for a Bot Service
 *
 * @summary creates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/PutEmailChannel.json
 */
async function createEmailChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.create(
    "OneResourceGroupName",
    "samplebotname",
    "EmailChannel",
    {
      location: "global",
      properties: {
        channelName: "EmailChannel",
        properties: {
          authMethod: 1,
          emailAddress: "a@b.com",
          isEnabled: true,
          magicCode: "000000",
        },
      },
    },
  );
  console.log(result);
}
