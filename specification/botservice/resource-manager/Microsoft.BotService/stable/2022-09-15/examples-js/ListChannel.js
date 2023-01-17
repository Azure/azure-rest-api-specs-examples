const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a Channel registration for a Bot Service including secrets
 *
 * @summary Lists a Channel registration for a Bot Service including secrets
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListChannel.json
 */
async function listChannel() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "EmailChannel";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.channels.listWithKeys(resourceGroupName, resourceName, channelName);
  console.log(result);
}
