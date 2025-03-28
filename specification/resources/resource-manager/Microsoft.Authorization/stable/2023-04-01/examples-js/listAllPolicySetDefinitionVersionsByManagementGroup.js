const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation lists all the policy set definition versions for all policy set definitions at the management group scope.
 *
 * @summary This operation lists all the policy set definition versions for all policy set definitions at the management group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listAllPolicySetDefinitionVersionsByManagementGroup.json
 */
async function listAllPolicyDefinitionVersionsAtManagementGroup() {
  const managementGroupName = "MyManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result =
    await client.policySetDefinitionVersions.listAllAtManagementGroup(managementGroupName);
  console.log(result);
}
