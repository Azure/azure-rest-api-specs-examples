const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation creates or updates a policy set definition in the given management group with the given name.
 *
 * @summary This operation creates or updates a policy set definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/createOrUpdatePolicySetDefinitionWithGroupsAtManagementGroup.json
 */
async function createOrUpdateAPolicySetDefinitionWithGroupsAtManagementGroupLevel() {
  const managementGroupId = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const parameters = {
    description: "Policies to enforce low cost storage SKUs",
    displayName: "Cost Management",
    metadata: { category: "Cost Management" },
    policyDefinitionGroups: [
      {
        name: "CostSaving",
        description: "Policies designed to control spend within a subscription.",
        displayName: "Cost Management Policies",
      },
      {
        name: "Organizational",
        description:
          "Policies that help enforce resource organization standards within a subscription.",
        displayName: "Organizational Policies",
      },
    ],
    policyDefinitions: [
      {
        groupNames: ["CostSaving"],
        parameters: {
          listOfAllowedSKUs: { value: ["Standard_GRS", "Standard_LRS"] },
        },
        policyDefinitionId:
          "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1",
        policyDefinitionReferenceId: "Limit_Skus",
      },
      {
        groupNames: ["Organizational"],
        parameters: { prefix: { value: "DeptA" }, suffix: { value: "-LC" } },
        policyDefinitionId:
          "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming",
        policyDefinitionReferenceId: "Resource_Naming",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitions.createOrUpdateAtManagementGroup(
    managementGroupId,
    policySetDefinitionName,
    parameters,
  );
  console.log(result);
}
