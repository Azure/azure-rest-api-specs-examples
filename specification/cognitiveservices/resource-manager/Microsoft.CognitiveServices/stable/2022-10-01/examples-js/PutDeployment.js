const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified deployments associated with the Cognitive Services account.
 *
 * @summary Update the state of specified deployments associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/PutDeployment.json
 */
async function putDeployment() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
  const accountName = "accountName";
  const deploymentName = "deploymentName";
  const deployment = {
    properties: {
      model: { name: "ada", format: "OpenAI", version: "1" },
      scaleSettings: { capacity: 1, scaleType: "Manual" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    deploymentName,
    deployment
  );
  console.log(result);
}

putDeployment().catch(console.error);
