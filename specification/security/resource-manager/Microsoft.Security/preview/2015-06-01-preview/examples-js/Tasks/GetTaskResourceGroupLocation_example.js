const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Recommended tasks that will help improve the security of the subscription proactively
 *
 * @summary Recommended tasks that will help improve the security of the subscription proactively
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTaskResourceGroupLocation_example.json
 */
async function getSecurityRecommendationTaskInAResourceGroup() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "myRg";
  const ascLocation = "westeurope";
  const taskName = "d55b4dc0-779c-c66c-33e5-d7bce24c4222";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.tasks.getResourceGroupLevelTask(
    resourceGroupName,
    ascLocation,
    taskName
  );
  console.log(result);
}

getSecurityRecommendationTaskInAResourceGroup().catch(console.error);
