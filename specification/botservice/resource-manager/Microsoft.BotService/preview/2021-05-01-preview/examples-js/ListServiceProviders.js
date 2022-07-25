const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available Service Providers for creating Connection Settings
 *
 * @summary Lists the available Service Providers for creating Connection Settings
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListServiceProviders.json
 */
async function listAuthServiceProviders() {
  const subscriptionId = "subscription-id";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.botConnection.listServiceProviders();
  console.log(result);
}

listAuthServiceProviders().catch(console.error);
