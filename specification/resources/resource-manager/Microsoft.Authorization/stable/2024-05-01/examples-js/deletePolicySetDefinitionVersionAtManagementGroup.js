const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes the policy set definition version in the given management group with the given name and version.
 *
 * @summary This operation deletes the policy set definition version in the given management group with the given name and version.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/deletePolicySetDefinitionVersionAtManagementGroup.json
 */
async function deleteAPolicySetDefinitionVersionAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const policyDefinitionVersion = "1.2.1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitionVersions.deleteAtManagementGroup(
    managementGroupName,
    policySetDefinitionName,
    policyDefinitionVersion,
  );
  console.log(result);
}
