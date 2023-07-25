const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified commitmentPlan associated with the Cognitive Services account.
 *
 * @summary Deletes the specified commitmentPlan associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/DeleteCommitmentPlan.json
 */
async function deleteCommitmentPlan() {
  const subscriptionId = process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["COGNITIVESERVICES_RESOURCE_GROUP"] || "resourceGroupName";
  const accountName = "accountName";
  const commitmentPlanName = "commitmentPlanName";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.commitmentPlans.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    commitmentPlanName
  );
  console.log(result);
}
