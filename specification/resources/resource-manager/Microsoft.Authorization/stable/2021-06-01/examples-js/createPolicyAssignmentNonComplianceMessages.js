const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to  This operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 *
 * @summary  This operation creates or updates a policy assignment with the given scope and name. Policy assignments apply to all resources contained within their scope. For example, when you assign a policy at resource group scope, that policy applies to all resources in the group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createPolicyAssignmentNonComplianceMessages.json
 */
async function createOrUpdateAPolicyAssignmentWithMultipleNonComplianceMessages() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const policyAssignmentName = "securityInitAssignment";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyAssignments.create(scope, policyAssignmentName, parameters);
  console.log(result);
}
