const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsDetailsOfTheAzureService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.services.get(serviceName);
  console.log(result);
}

getsDetailsOfTheAzureService().catch(console.error);
