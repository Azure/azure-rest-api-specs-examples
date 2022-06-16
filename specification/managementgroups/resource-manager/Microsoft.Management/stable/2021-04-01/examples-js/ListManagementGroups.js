const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List management groups for the authenticated user.

 *
 * @summary List management groups for the authenticated user.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListManagementGroups.json
 */
async function listManagementGroups() {
  const cacheControl = "no-cache";
  const options = { cacheControl };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const resArray = new Array();
  for await (let item of client.managementGroups.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagementGroups().catch(console.error);
