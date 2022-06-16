const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables the activation of a snoozed or dismissed recommendation. The snoozed or dismissed attribute of a recommendation is referred to as a suppression.
 *
 * @summary Enables the activation of a snoozed or dismissed recommendation. The snoozed or dismissed attribute of a recommendation is referred to as a suppression.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/DeleteSuppression.json
 */
async function deleteSuppression() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri = "resourceUri";
  const recommendationId = "recommendationId";
  const name = "suppressionName1";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.suppressions.delete(resourceUri, recommendationId, name);
  console.log(result);
}

deleteSuppression().catch(console.error);
