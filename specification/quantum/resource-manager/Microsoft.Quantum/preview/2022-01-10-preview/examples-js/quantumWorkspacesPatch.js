const { AzureQuantumManagementClient } = require("@azure/arm-quantum");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing workspace's tags.
 *
 * @summary Updates an existing workspace's tags.
 * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesPatch.json
 */
async function quantumWorkspacesPatchTags() {
  const subscriptionId =
    process.env["QUANTUM_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["QUANTUM_RESOURCE_GROUP"] || "quantumResourcegroup";
  const workspaceName = "quantumworkspace1";
  const workspaceTags = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureQuantumManagementClient(credential, subscriptionId);
  const result = await client.workspaces.updateTags(
    resourceGroupName,
    workspaceName,
    workspaceTags
  );
  console.log(result);
}
