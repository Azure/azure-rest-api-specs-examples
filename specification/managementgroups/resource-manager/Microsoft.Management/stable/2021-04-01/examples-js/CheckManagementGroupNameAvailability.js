const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if the specified management group name is valid and unique
 *
 * @summary Checks if the specified management group name is valid and unique
 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/CheckManagementGroupNameAvailability.json
 */
async function checkManagementGroupNameAvailability() {
  const checkNameAvailabilityRequest = {
    name: "nameTocheck",
    type: "Microsoft.Management/managementGroups",
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.checkNameAvailability(checkNameAvailabilityRequest);
  console.log(result);
}

checkManagementGroupNameAvailability().catch(console.error);
