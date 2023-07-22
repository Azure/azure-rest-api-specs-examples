const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get role definition by ID (GUID).
 *
 * @summary Get role definition by ID (GUID).
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-05-01-preview/examples/GetRoleDefinitionById.json
 */
async function getRoleDefinitionById() {
  const scope = "scope";
  const roleDefinitionId = "roleDefinitionId";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleDefinitions.get(scope, roleDefinitionId);
  console.log(result);
}
