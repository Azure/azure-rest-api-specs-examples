const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List ModelCapacities.
 *
 * @summary List ModelCapacities.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ListModelCapacities.json
 */
async function listModelCapacities() {
  const subscriptionId =
    process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const modelFormat = "OpenAI";
  const modelName = "ada";
  const modelVersion = "1";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.modelCapacities.list(modelFormat, modelName, modelVersion)) {
    resArray.push(item);
  }
  console.log(resArray);
}
