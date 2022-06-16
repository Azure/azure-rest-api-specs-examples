const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details of the management group.

 *
 * @summary Get the details of the management group.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithExpandAndRecurse.json
 */
async function getManagementGroupsWithExpandAndRecurse() {
  const groupId = "20000000-0001-0000-0000-000000000000";
  const expand = "children";
  const recurse = true;
  const cacheControl = "no-cache";
  const options = {
    expand,
    recurse,
    cacheControl,
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.managementGroups.get(groupId, options);
  console.log(result);
}

getManagementGroupsWithExpandAndRecurse().catch(console.error);
