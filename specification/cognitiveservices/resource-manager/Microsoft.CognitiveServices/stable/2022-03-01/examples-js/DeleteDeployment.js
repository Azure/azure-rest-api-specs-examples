const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified deployment associated with the Cognitive Services account.
 *
 * @summary Deletes the specified deployment associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/DeleteDeployment.json
 */
async function deleteDeployment() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
  const accountName = "accountName";
  const deploymentName = "deploymentName";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    deploymentName
  );
  console.log(result);
}

deleteDeployment().catch(console.error);
