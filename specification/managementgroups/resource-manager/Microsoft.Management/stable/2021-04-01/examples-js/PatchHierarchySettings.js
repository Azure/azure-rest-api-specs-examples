const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the hierarchy settings defined at the Management Group level.

 *
 * @summary Updates the hierarchy settings defined at the Management Group level.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PatchHierarchySettings.json
 */
async function getGroupSettings() {
  const groupId = "root";
  const createTenantSettingsRequest = {
    defaultManagementGroup: "/providers/Microsoft.Management/managementGroups/DefaultGroup",
    requireAuthorizationForGroupCreation: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.hierarchySettingsOperations.update(
    groupId,
    createTenantSettingsRequest
  );
  console.log(result);
}

getGroupSettings().catch(console.error);
