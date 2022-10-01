const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add existing user to existing group
 *
 * @summary Add existing user to existing group
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroupUser.json
 */
async function apiManagementCreateGroupUser() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "tempgroup";
  const userId = "59307d350af58404d8a26300";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.groupUser.create(resourceGroupName, serviceName, groupId, userId);
  console.log(result);
}

apiManagementCreateGroupUser().catch(console.error);
