const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation deletes a variable value, given its name, the management group it was created in, and the variable it belongs to. The scope of a variable value is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 *
 * @summary This operation deletes a variable value, given its name, the management group it was created in, and the variable it belongs to. The scope of a variable value is the part of its ID preceding '/providers/Microsoft.Authorization/variables/{variableName}'.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/deleteVariableValueAtManagementGroup.json
 */
async function deleteAVariableValueAtManagementGroup() {
  const managementGroupId = "DevOrg";
  const variableName = "DemoTestVariable";
  const variableValueName = "TestValue";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.variableValues.deleteAtManagementGroup(
    managementGroupId,
    variableName,
    variableValueName,
  );
  console.log(result);
}
