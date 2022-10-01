const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of user entities associated with the group.
 *
 * @summary Lists a collection of user entities associated with the group.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListGroupUsers.json
 */
async function apiManagementListGroupUsers() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "57d2ef278aa04f0888cba3f3";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.groupUser.list(resourceGroupName, serviceName, groupId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListGroupUsers().catch(console.error);
