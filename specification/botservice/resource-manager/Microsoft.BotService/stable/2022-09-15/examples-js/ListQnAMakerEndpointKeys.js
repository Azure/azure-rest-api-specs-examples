const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the QnA Maker endpoint keys
 *
 * @summary Lists the QnA Maker endpoint keys
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListQnAMakerEndpointKeys.json
 */
async function listQnAMakerEndpointKeys() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subscription-id";
  const parameters = {
    authkey: "testAuthKey",
    hostname: "https://xxx.cognitiveservices.azure.com/",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.qnAMakerEndpointKeys.get(parameters);
  console.log(result);
}
