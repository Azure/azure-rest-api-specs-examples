const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a role definition.
 *
 * @summary Creates or updates a role definition.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/PutRoleDefinition.json
 */
async function createRoleDefinition() {
  const subscriptionId =
    process.env["AUTHORIZATION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "scope";
  const roleDefinitionId = "roleDefinitionId";
  const roleDefinition = {};
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential, subscriptionId);
  const result = await client.roleDefinitions.createOrUpdate(
    scope,
    roleDefinitionId,
    roleDefinition
  );
  console.log(result);
}
