const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a Channel registration from a Bot Service
 *
 * @summary deletes a Channel registration from a Bot Service
 * x-ms-original-file: 2023-09-15-preview/DeleteChannel.json
 */
async function deleteChannel() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  await client.channels.delete("OneResourceGroupName", "samplebotname", "EmailChannel");
}
