const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the hierarchy settings defined at the Management Group level.

 *
 * @summary Deletes the hierarchy settings defined at the Management Group level.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/DeleteHierarchySettings.json
 */
async function getGroupSettings() {
  const groupId = "root";
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.hierarchySettingsOperations.delete(groupId);
  console.log(result);
}

getGroupSettings().catch(console.error);
