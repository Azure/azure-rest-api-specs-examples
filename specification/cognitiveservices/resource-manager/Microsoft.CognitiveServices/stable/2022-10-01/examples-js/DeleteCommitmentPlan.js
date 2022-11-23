const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified commitmentPlan associated with the Cognitive Services account.
 *
 * @summary Deletes the specified commitmentPlan associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeleteCommitmentPlan.json
 */
async function deleteCommitmentPlan() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
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

deleteCommitmentPlan().catch(console.error);
