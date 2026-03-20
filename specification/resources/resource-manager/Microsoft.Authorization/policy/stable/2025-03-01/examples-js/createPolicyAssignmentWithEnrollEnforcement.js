const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary this operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: 2025-03-01/createPolicyAssignmentWithEnrollEnforcement.json
 */
async function createOrUpdateAPolicyAssignmentToEnforcePolicyEffectOnlyOnEnrolledResourcesDuringResourceCreationOrUpdate() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyAssignments.create(
    "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
    "EnforceNamingEnroll",
    {
      description: "Force resource names to begin with given DeptA and end with -LC",
      displayName: "Enforce resource naming rules",
      enforcementMode: "Enroll",
      metadata: { assignedBy: "Special Someone" },
      parameters: { prefix: { value: "DeptA" }, suffix: { value: "-LC" } },
      policyDefinitionId:
        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming",
    },
  );
  console.log(result);
}
