const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the solution list. It will retrieve both first party and third party solutions
 *
 * @summary Retrieves the solution list. It will retrieve both first party and third party solutions
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionListForSubscription.json
 */
async function solutionList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.solutions.listBySubscription();
  console.log(result);
}

solutionList().catch(console.error);
