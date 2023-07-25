const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Commitment Tiers.
 *
 * @summary List Commitment Tiers.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListCommitmentTiers.json
 */
async function listCommitmentTiers() {
  const subscriptionId = process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "subscriptionId";
  const location = "location";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.commitmentTiers.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
