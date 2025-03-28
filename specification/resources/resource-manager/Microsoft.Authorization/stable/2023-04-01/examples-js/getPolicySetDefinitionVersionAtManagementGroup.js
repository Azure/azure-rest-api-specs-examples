const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the policy set definition version in the given management group with the given name and version.
 *
 * @summary This operation retrieves the policy set definition version in the given management group with the given name and version.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getPolicySetDefinitionVersionAtManagementGroup.json
 */
async function retrieveAPolicySetDefinitionVersionAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const policyDefinitionVersion = "1.2.1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitionVersions.getAtManagementGroup(
    managementGroupName,
    policySetDefinitionName,
    policyDefinitionVersion,
  );
  console.log(result);
}
