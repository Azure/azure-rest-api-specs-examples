const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the Solution.
 *
 * @summary Creates or updates the Solution.
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionCreate.json
 */
async function solutionCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const solutionName = "solution1";
  const parameters = {
    location: "East US",
    plan: {
      name: "name1",
      product: "product1",
      promotionCode: "promocode1",
      publisher: "publisher1",
    },
    properties: {
      containedResources: [
        "/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1",
        "/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2",
      ],
      referencedResources: [
        "/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2",
        "/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3",
      ],
      workspaceResourceId:
        "/subscriptions/sub2/resourceGroups/rg2/providers/Microsoft.OperationalInsights/workspaces/ws1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.solutions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    solutionName,
    parameters
  );
  console.log(result);
}

solutionCreate().catch(console.error);
