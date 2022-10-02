const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all user groups.
 *
 * @summary Lists all user groups.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListUserGroups.json
 */
async function apiManagementListUserGroups() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const userId = "57681833a40f7eb6c49f6acf";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.userGroup.list(resourceGroupName, serviceName, userId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListUserGroups().catch(console.error);
