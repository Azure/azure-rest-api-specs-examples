const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Channel registration for a Bot Service
 *
 * @summary Creates a Channel registration for a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutChannel.json
 */
async function createBot() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "EmailChannel";
  const parameters = {
    location: "global",
    properties: {
      channelName: "EmailChannel",
      properties: { emailAddress: "a@b.com", isEnabled: true, password: "pwd" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.create(
    resourceGroupName,
    resourceName,
    channelName,
    parameters
  );
  console.log(result);
}

createBot().catch(console.error);
