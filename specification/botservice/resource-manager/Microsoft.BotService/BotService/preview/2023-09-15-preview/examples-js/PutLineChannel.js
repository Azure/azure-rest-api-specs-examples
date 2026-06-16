const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Channel registration for a Bot Service
 *
 * @summary creates a Channel registration for a Bot Service
 * x-ms-original-file: 2023-09-15-preview/PutLineChannel.json
 */
async function createLineChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.create(
    "OneResourceGroupName",
    "samplebotname",
    "LineChannel",
    {
      location: "global",
      properties: {
        channelName: "LineChannel",
        properties: {
          lineRegistrations: [
            { channelAccessToken: "channelAccessToken", channelSecret: "channelSecret" },
          ],
        },
      },
    },
  );
  console.log(result);
}
