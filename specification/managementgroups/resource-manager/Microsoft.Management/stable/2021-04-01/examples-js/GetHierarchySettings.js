const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management Group of the hierarchy.

 *
 * @summary Gets the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management Group of the hierarchy.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetHierarchySettings.json
 */
async function getGroupSettings() {
  const groupId = "root";
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.hierarchySettingsOperations.get(groupId);
  console.log(result);
}

getGroupSettings().catch(console.error);
