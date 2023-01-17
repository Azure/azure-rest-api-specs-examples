const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all the resources of a particular type belonging to a resource group
 *
 * @summary Returns all the resources of a particular type belonging to a resource group
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListBotsByResourceGroup.json
 */
async function listBotsByResourceGroup() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "OneResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bots.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
