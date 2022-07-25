const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get per subscription settings needed to host bot in compute resource such as Azure App Service
 *
 * @summary Get per subscription settings needed to host bot in compute resource such as Azure App Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetHostSettings.json
 */
async function getBotHostSettings() {
  const subscriptionId = "subscription-id";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.hostSettings.get();
  console.log(result);
}

getBotHostSettings().catch(console.error);
