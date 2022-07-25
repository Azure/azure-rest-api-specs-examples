const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Channel registration for a Bot Service
 *
 * @summary Updates a Channel registration for a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateChannel.json
 */
async function updateBot() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "EmailChannel";
  const location = "global";
  const properties = {
    channelName: "EmailChannel",
    properties: { emailAddress: "a@b.com", isEnabled: true, password: "pwd" },
  };
  const options = { location, properties };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.update(
    resourceGroupName,
    resourceName,
    channelName,
    options
  );
  console.log(result);
}

updateBot().catch(console.error);
