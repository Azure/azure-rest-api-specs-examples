const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all entities that descend from a management group.

 *
 * @summary List all entities that descend from a management group.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetDescendants.json
 */
async function getDescendants() {
  const groupId = "20000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const resArray = new Array();
  for await (let item of client.managementGroups.listDescendants(groupId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDescendants().catch(console.error);
