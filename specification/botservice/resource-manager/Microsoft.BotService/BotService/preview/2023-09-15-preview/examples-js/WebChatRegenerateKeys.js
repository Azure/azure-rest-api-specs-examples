const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to regenerates secret keys and returns them for the DirectLine Channel of a particular BotService resource
 *
 * @summary regenerates secret keys and returns them for the DirectLine Channel of a particular BotService resource
 * x-ms-original-file: 2023-09-15-preview/WebChatRegenerateKeys.json
 */
async function regenerateKeysForWebChatChannelSite() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscription-id";
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.directLine.regenerateKeys(
    "OneResourceGroupName",
    "samplebotname",
    "WebChatChannel",
    { key: "key1", siteName: "testSiteName" },
  );
  console.log(result);
}
