const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a Bot.
 *
 * @summary Gets the private link resources that need to be created for a Bot.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListPrivateLinkResources.json
 */
async function listPrivateLinkResources() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "res6977";
  const resourceName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.privateLinkResources.listByBotResource(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}
