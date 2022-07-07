const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get problem classification details for a specific Azure service.
 *
 * @summary Get problem classification details for a specific Azure service.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetProblemClassification.json
 */
async function getsDetailsOfProblemClassificationForAzureService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const problemClassificationName = "problemClassification_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.problemClassifications.get(serviceName, problemClassificationName);
  console.log(result);
}

getsDetailsOfProblemClassificationForAzureService().catch(console.error);
