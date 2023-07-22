const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a role definition by ID.
 *
 * @summary Gets a role definition by ID.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-05-01-preview/examples/GetRoleDefinitionByRoleId.json
 */
async function getRoleDefinitionByRoleId() {
  const roleId = "roleDefinitionId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleDefinitions.getById(roleId);
  console.log(result);
}
