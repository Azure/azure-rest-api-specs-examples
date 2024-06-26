const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the solution in the subscription.
 *
 * @summary Deletes the solution in the subscription.
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionDelete.json
 */
async function solutionDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const solutionName = "solution1";
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.solutions.beginDeleteAndWait(resourceGroupName, solutionName);
  console.log(result);
}

solutionDelete().catch(console.error);
