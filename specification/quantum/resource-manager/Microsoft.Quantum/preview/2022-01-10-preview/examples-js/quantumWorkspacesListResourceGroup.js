const { AzureQuantumManagementClient } = require("@azure/arm-quantum");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Workspaces within a resource group.
 *
 * @summary Gets the list of Workspaces within a resource group.
 * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesListResourceGroup.json
 */
async function quantumWorkspacesListByResourceGroup() {
  const subscriptionId =
    process.env["QUANTUM_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["QUANTUM_RESOURCE_GROUP"] || "quantumResourcegroup";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuantumManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
