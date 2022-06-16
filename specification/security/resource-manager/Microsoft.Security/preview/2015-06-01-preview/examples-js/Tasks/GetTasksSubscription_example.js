const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Recommended tasks that will help improve the security of the subscription proactively
 *
 * @summary Recommended tasks that will help improve the security of the subscription proactively
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTasksSubscription_example.json
 */
async function getSecurityRecommendationsTasks() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tasks.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getSecurityRecommendationsTasks().catch(console.error);
