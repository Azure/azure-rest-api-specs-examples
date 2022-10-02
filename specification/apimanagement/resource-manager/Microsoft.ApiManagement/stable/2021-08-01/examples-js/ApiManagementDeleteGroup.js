const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes specific group of the API Management service instance.
 *
 * @summary Deletes specific group of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteGroup.json
 */
async function apiManagementDeleteGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "aadGroup";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.group.delete(resourceGroupName, serviceName, groupId, ifMatch);
  console.log(result);
}

apiManagementDeleteGroup().catch(console.error);
