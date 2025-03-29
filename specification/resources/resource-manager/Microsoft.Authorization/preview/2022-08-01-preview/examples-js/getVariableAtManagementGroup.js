const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves a single variable, given its name and the  management group it was created at.
 *
 * @summary This operation retrieves a single variable, given its name and the  management group it was created at.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/getVariableAtManagementGroup.json
 */
async function retrieveAVariableAtManagementGroup() {
  const managementGroupId = "DevOrg";
  const variableName = "DemoTestVariable";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.variables.getAtManagementGroup(managementGroupId, variableName);
  console.log(result);
}
