const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the commitmentPlans associated with the Cognitive Services account.
 *
 * @summary Gets the commitmentPlans associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListCommitmentPlans.json
 */
async function listCommitmentPlans() {
  const subscriptionId = process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["COGNITIVESERVICES_RESOURCE_GROUP"] || "resourceGroupName";
  const accountName = "accountName";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.commitmentPlans.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
