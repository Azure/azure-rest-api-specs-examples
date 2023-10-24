const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific Azure service for support ticket creation.
 *
 * @summary Gets a specific Azure service for support ticket creation.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetService.json
 */
async function getsDetailsOfTheAzureService() {
  const serviceName = "service_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.services.get(serviceName);
  console.log(result);
}
