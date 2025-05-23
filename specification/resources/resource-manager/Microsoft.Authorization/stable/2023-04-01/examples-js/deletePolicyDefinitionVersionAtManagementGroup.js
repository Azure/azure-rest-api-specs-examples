const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes the policy definition in the given management group with the given name.
 *
 * @summary This operation deletes the policy definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/deletePolicyDefinitionVersionAtManagementGroup.json
 */
async function deleteAPolicyDefinitionVersionAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policyDefinitionName = "ResourceNaming";
  const policyDefinitionVersion = "1.2.1";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyDefinitionVersions.deleteAtManagementGroup(
    managementGroupName,
    policyDefinitionName,
    policyDefinitionVersion,
  );
  console.log(result);
}
