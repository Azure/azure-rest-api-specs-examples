const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Channel registration for a Bot Service
 *
 * @summary updates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/UpdateLineChannel.json
 */
async function updateLineChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.update(
    "OneResourceGroupName",
    "samplebotname",
    "LineChannel",
    {
      properties: {
        channelName: "LineChannel",
        properties: {
          lineRegistrations: [
            { channelAccessToken: "channelAccessToken", channelSecret: "channelSecret" },
          ],
        },
      },
      location: "global",
    },
  );
  console.log(result);
}
