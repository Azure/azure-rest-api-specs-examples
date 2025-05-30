const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a list of all the policy definition versions for the given policy definition in the given management group.
 *
 * @summary This operation retrieves a list of all the policy definition versions for the given policy definition in the given management group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2024-05-01/examples/listPolicyDefinitionVersionsByManagementGroup.json
 */
async function listPolicyDefinitionVersionsByManagementGroup() {
  const managementGroupName = "MyManagementGroup";
  const policyDefinitionName = "ResourceNaming";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const resArray = new Array();
  for await (const item of client.policyDefinitionVersions.listByManagementGroup(
    managementGroupName,
    policyDefinitionName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
