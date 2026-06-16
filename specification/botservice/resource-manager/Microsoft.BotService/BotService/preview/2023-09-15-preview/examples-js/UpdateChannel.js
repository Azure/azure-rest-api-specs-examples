const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Channel registration for a Bot Service
 *
 * @summary updates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/UpdateChannel.json
 */
async function updateChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.update(
    "OneResourceGroupName",
    "samplebotname",
    "EmailChannel",
    {
      properties: {
        channelName: "EmailChannel",
        properties: { emailAddress: "a@b.com", isEnabled: true, password: "pwd" },
      },
      location: "global",
    },
  );
  console.log(result);
}
