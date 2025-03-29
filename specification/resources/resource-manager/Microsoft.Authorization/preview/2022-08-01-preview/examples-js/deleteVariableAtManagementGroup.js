const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes a variable, given its name and the management group it was created in. The scope of a variable is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 *
 * @summary This operation deletes a variable, given its name and the management group it was created in. The scope of a variable is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/deleteVariableAtManagementGroup.json
 */
async function deleteAVariableAtManagementGroup() {
  const managementGroupId = "DevOrg";
  const variableName = "DemoTestVariable";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.variables.deleteAtManagementGroup(managementGroupId, variableName);
  console.log(result);
}
