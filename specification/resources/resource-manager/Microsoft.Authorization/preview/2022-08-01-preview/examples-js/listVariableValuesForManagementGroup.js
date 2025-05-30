const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the list of all variable values applicable the variable indicated at the management group scope.
 *
 * @summary This operation retrieves the list of all variable values applicable the variable indicated at the management group scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/listVariableValuesForManagementGroup.json
 */
async function listVariableValuesAtAManagementGroupScope() {
  const managementGroupId = "DevOrg";
  const variableName = "DemoTestVariable";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const resArray = new Array();
  for await (const item of client.variableValues.listForManagementGroup(
    managementGroupId,
    variableName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
