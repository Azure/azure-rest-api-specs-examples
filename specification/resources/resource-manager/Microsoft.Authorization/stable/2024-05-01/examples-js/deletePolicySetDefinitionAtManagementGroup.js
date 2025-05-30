const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes the policy set definition in the given management group with the given name.
 *
 * @summary This operation deletes the policy set definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/deletePolicySetDefinitionAtManagementGroup.json
 */
async function deleteAPolicySetDefinitionAtManagementGroupLevel() {
  const managementGroupId = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitions.deleteAtManagementGroup(
    managementGroupId,
    policySetDefinitionName,
  );
  console.log(result);
}
