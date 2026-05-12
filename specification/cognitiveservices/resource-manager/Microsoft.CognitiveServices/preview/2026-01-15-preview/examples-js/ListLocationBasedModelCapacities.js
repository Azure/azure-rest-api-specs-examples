const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Location Based ModelCapacities.
 *
 * @summary list Location Based ModelCapacities.
 * x-ms-original-file: 2026-01-15-preview/ListLocationBasedModelCapacities.json
 */
async function listLocationBasedModelCapacities() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.locationBasedModelCapacities.list(
    "WestUS",
    "OpenAI",
    "ada",
    "1",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
