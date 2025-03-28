const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the policy set definition in the given management group with the given name.
 *
 * @summary This operation retrieves the policy set definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getPolicySetDefinitionAtManagementGroup.json
 */
async function retrieveAPolicySetDefinitionAtManagementGroupLevel() {
  const managementGroupId = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policySetDefinitions.getAtManagementGroup(
    managementGroupId,
    policySetDefinitionName,
  );
  console.log(result);
}
