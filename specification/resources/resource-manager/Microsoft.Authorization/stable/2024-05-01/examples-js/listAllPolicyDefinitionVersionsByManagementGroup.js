const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation lists all the policy definition versions for all policy definitions at the management group scope.
 *
 * @summary This operation lists all the policy definition versions for all policy definitions at the management group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/listAllPolicyDefinitionVersionsByManagementGroup.json
 */
async function listAllPolicyDefinitionVersionsAtManagementGroup() {
  const managementGroupName = "MyManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result =
    await client.policyDefinitionVersions.listAllAtManagementGroup(managementGroupName);
  console.log(result);
}
