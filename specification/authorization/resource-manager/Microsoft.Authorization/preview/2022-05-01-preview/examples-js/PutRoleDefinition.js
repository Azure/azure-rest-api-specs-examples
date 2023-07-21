const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a role definition.
 *
 * @summary Creates or updates a role definition.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-05-01-preview/examples/PutRoleDefinition.json
 */
async function createRoleDefinition() {
  const scope = "scope";
  const roleDefinitionId = "roleDefinitionId";
  const roleDefinition = {};
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleDefinitions.createOrUpdate(
    scope,
    roleDefinitionId,
    roleDefinition
  );
  console.log(result);
}
