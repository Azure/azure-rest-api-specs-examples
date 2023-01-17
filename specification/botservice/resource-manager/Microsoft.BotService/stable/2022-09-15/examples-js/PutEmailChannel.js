const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Channel registration for a Bot Service
 *
 * @summary Creates a Channel registration for a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/PutEmailChannel.json
 */
async function createEmailChannel() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "EmailChannel";
  const parameters = {
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
