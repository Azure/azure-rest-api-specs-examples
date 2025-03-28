const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a list of all the policy set definition versions for the given policy set definition in a given management group.
 *
 * @summary This operation retrieves a list of all the policy set definition versions for the given policy set definition in a given management group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listPolicySetDefinitionVersionsByManagementGroup.json
 */
async function listPolicySetDefinitionsAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policySetDefinitionName = "CostManagement";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const resArray = new Array();
  for await (const item of client.policySetDefinitionVersions.listByManagementGroup(
    managementGroupName,
    policySetDefinitionName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
