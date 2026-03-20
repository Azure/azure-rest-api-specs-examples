const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2025-03-01/createPolicyAssignmentNonComplianceMessages.json
 */
async function createOrUpdateAPolicyAssignmentWithMultipleNonComplianceMessages() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.create(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "securityInitAssignment",
    {
      displayName: "Enforce security policies",
      nonComplianceMessages: [
        {
          message:
            "Resources must comply with all internal security policies. See <internal site URL> for more info.",
        },
        {
          message: "Resource names must start with 'DeptA' and end with '-LC'.",
          policyDefinitionReferenceId: "10420126870854049575",
        },
        {
          message: "Storage accounts must have firewall rules configured.",
          policyDefinitionReferenceId: "8572513655450389710",
        },
      ],
      policyDefinitionId:
        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/securityInitiative",
    },
  );
  console.log(result);
}
