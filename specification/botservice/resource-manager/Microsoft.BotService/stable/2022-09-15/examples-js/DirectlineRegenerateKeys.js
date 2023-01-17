const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates secret keys and returns them for the DirectLine Channel of a particular BotService resource
 *
 * @summary Regenerates secret keys and returns them for the DirectLine Channel of a particular BotService resource
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DirectlineRegenerateKeys.json
 */
async function regenerateKeysForDirectLineChannelSite() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "OneResourceGroupName";
  const resourceName = "samplebotname";
  const channelName = "DirectLineChannel";
  const parameters = { key: "key1", siteName: "testSiteName" };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.directLine.regenerateKeys(
    resourceGroupName,
    resourceName,
    channelName,
    parameters
  );
  console.log(result);
}
