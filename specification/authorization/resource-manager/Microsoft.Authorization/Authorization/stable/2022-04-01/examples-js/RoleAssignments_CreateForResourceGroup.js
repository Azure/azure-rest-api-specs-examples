const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a role assignment by scope and name.
 *
 * @summary create or update a role assignment by scope and name.
 * x-ms-original-file: 2022-04-01/RoleAssignments_CreateForResourceGroup.json
 */
async function createRoleAssignmentForResourceGroup() {
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.roleAssignments.create(
    "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg",
    "05c5a614-a7d6-4502-b150-c2fb455033ff",
    {
      principalId: "ce2ce14e-85d7-4629-bdbc-454d0519d987",
      principalType: "User",
      roleDefinitionId:
        "/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d",
    },
  );
  console.log(result);
}
