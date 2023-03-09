const { MicrosoftVoiceServices } = require("@azure/arm-voiceservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List CommunicationsGateway resources by resource group
 *
 * @summary List CommunicationsGateway resources by resource group
 * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_ListByResourceGroup.json
 */
async function listCommunicationsGatewayResource() {
  const subscriptionId =
    process.env["VOICESERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["VOICESERVICES_RESOURCE_GROUP"] || "testrg";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftVoiceServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationsGateways.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
