const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the solution list. It will retrieve both first party and third party solutions
 *
 * @summary Retrieves the solution list. It will retrieve both first party and third party solutions
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionList.json
 */
async function solutionList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.solutions.listByResourceGroup(resourceGroupName);
  console.log(result);
}

solutionList().catch(console.error);
