const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Connection Setting registration for a Bot Service
 *
 * @summary Get a Connection Setting registration for a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetConnection.json
 */
async function updateConnectionSetting() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const connectionName = "sampleConnection";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.botConnection.get(resourceGroupName, resourceName, connectionName);
  console.log(result);
}

updateConnectionSetting().catch(console.error);
