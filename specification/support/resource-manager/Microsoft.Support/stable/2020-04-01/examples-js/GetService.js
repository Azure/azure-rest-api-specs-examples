const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific Azure service for support ticket creation.
 *
 * @summary Gets a specific Azure service for support ticket creation.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetService.json
 */
async function getsDetailsOfTheAzureService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.services.get(serviceName);
  console.log(result);
}

getsDetailsOfTheAzureService().catch(console.error);
