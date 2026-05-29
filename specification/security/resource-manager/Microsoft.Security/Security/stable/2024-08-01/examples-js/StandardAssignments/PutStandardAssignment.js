const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a standard assignment with the given scope and name. standard assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation creates or updates a standard assignment with the given scope and name. standard assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2024-08-01/StandardAssignments/PutStandardAssignment.json
 */
async function putAnAuditStandardAssignment() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.standardAssignments.create(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
    "1f3afdf9-d0c9-4c3d-847f-89da613e70a8",
    {
      description: "Set of policies monitored by Azure Security Center for cross cloud",
      assignedStandard: {
        id: "/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8",
      },
      displayName: "ASC Default",
      effect: "Audit",
      excludedScopes: [],
    },
  );
  console.log(result);
}
