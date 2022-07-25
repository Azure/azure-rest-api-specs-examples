const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check whether a workspace name is available
 *
 * @summary Check whether a workspace name is available
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CheckNameAvailabilityWorkspaceAvailable.json
 */
async function checkForAWorkspaceNameThatIsAvailable() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const request = {
    name: "workspace1",
    type: "Microsoft.ProjectArcadia/workspaces",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.operations.checkNameAvailability(request);
  console.log(result);
}

checkForAWorkspaceNameThatIsAvailable().catch(console.error);
