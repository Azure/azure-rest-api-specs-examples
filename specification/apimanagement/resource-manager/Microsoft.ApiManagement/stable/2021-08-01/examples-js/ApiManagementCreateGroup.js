const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a group.
 *
 * @summary Creates or Updates a group.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroup.json
 */
async function apiManagementCreateGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "tempgroup";
  const parameters = { displayName: "temp group" };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.group.createOrUpdate(
    resourceGroupName,
    serviceName,
    groupId,
    parameters
  );
  console.log(result);
}

apiManagementCreateGroup().catch(console.error);
