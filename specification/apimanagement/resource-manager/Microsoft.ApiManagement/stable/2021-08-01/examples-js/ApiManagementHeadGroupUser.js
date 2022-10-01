const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that user entity specified by identifier is associated with the group entity.
 *
 * @summary Checks that user entity specified by identifier is associated with the group entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadGroupUser.json
 */
async function apiManagementHeadGroupUser() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "59306a29e4bbd510dc24e5f9";
  const userId = "5931a75ae4bbd512a88c680b";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.groupUser.checkEntityExists(
    resourceGroupName,
    serviceName,
    groupId,
    userId
  );
  console.log(result);
}

apiManagementHeadGroupUser().catch(console.error);
