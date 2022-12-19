const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the principal assignment name is valid and is not already in use.
 *
 * @summary Checks that the principal assignment name is valid and is not already in use.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsCheckNameAvailability.json
 */
async function kustoPoolPrincipalAssignmentsCheckNameAvailability() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const principalAssignmentName = {
    name: "kustoprincipal1",
    type: "Microsoft.Synapse/workspaces/kustoPools/principalAssignments",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolPrincipalAssignments.checkNameAvailability(
    workspaceName,
    kustoPoolName,
    resourceGroupName,
    principalAssignmentName
  );
  console.log(result);
}
