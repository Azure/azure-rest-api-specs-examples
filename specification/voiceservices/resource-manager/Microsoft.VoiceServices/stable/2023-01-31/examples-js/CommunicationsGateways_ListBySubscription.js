const { MicrosoftVoiceServices } = require("@azure/arm-voiceservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List CommunicationsGateway resources by subscription ID
 *
 * @summary List CommunicationsGateway resources by subscription ID
 * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_ListBySubscription.json
 */
async function listCommunicationsGatewayResourceSub() {
  const subscriptionId =
    process.env["VOICESERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftVoiceServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationsGateways.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
