const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list ModelCapacities.
 *
 * @summary list ModelCapacities.
 * x-ms-original-file: 2026-01-15-preview/ListModelCapacitiesDataZoneScope.json
 */
async function listModelCapacitiesDataZoneScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.modelCapacities.list("OpenAI", "ada", "1")) {
    resArray.push(item);
  }

  console.log(resArray);
}
