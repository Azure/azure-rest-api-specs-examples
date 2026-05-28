const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the details of the management group.
 *
 * @summary get the details of the management group.
 * x-ms-original-file: 2023-04-01/GetManagementGroupWithExpand.json
 */
async function getManagementGroupWithExpand() {
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.managementGroups.get("20000000-0001-0000-0000-000000000000", {
    expand: "children",
  });
  console.log(result);
}
