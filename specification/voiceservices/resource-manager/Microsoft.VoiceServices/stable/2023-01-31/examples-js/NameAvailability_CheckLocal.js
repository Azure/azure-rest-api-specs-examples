const { MicrosoftVoiceServices } = require("@azure/arm-voiceservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check whether the resource name is available in the given region.
 *
 * @summary Check whether the resource name is available in the given region.
 * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/NameAvailability_CheckLocal.json
 */
async function checkLocalNameAvailability() {
  const subscriptionId =
    process.env["VOICESERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "useast";
  const body = {
    name: "myname",
    type: "Microsoft.VoiceServices/CommunicationsGateways",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftVoiceServices(credential, subscriptionId);
  const result = await client.nameAvailability.checkLocal(location, body);
  console.log(result);
}
