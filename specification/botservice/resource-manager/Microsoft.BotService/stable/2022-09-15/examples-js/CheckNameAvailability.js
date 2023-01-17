const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check whether a bot name is available.
 *
 * @summary Check whether a bot name is available.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/CheckNameAvailability.json
 */
async function checkNameAvailability() {
  const subscriptionId =
    process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const parameters = {
    name: "testbotname",
    type: "string",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.bots.getCheckNameAvailability(parameters);
  console.log(result);
}
