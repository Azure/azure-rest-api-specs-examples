const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the status of the recommendation computation or generation process. Invoke this API after calling the generation recommendation. The URI of this API is returned in the Location field of the response header.
 *
 * @summary Retrieves the status of the recommendation computation or generation process. Invoke this API after calling the generation recommendation. The URI of this API is returned in the Location field of the response header.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/EmptyResponse.json
 */
async function getGenerateStatus() {
  const subscriptionId = "subscriptionId";
  const operationId = "operationGUID";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.recommendations.getGenerateStatus(operationId);
  console.log(result);
}

getGenerateStatus().catch(console.error);
