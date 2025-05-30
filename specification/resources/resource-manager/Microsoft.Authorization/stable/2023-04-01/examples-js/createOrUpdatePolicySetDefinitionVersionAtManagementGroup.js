const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation creates or updates a policy set definition version in the given management group with the given name and version.
 *
 * @summary This operation creates or updates a policy set definition version in the given management group with the given name and version.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createOrUpdatePolicySetDefinitionVersionAtManagementGroup.json
 */
async function createOrUpdateAPolicySetDefinitionVersionAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const policyDefinitionVersion = "1.2.1";
  const parameters = {
    description: "Policies to enforce low cost storage SKUs",
    displayName: "Cost Management",
    metadata: { category: "Cost Management" },
    policyDefinitions: [
      {
        parameters: {
          listOfAllowedSKUs: { value: ["Standard_GRS", "Standard_LRS"] },
        },
        policyDefinitionId:
          "/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1",
        policyDefinitionReferenceId: "Limit_Skus",
      },
      {
        parameters: { prefix: { value: "DeptA" }, suffix: { value: "-LC" } },
        policyDefinitionId:
          "/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming",
        policyDefinitionReferenceId: "Resource_Naming",
      },
    ],
    version: "1.2.1",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitionVersions.createOrUpdateAtManagementGroup(
    managementGroupName,
    policySetDefinitionName,
    policyDefinitionVersion,
    parameters,
  );
  console.log(result);
}
