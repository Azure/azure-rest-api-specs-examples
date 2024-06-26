const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a Solution. Only updating tags supported.
 *
 * @summary Patch a Solution. Only updating tags supported.
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionUpdate.json
 */
async function solutionUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const solutionName = "solution1";
  const parameters = {
    tags: { dept: "IT", environment: "Test" },
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.solutions.beginUpdateAndWait(
    resourceGroupName,
    solutionName,
    parameters
  );
  console.log(result);
}

solutionUpdate().catch(console.error);
