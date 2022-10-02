const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes specific tag of the API Management service instance.
 *
 * @summary Deletes specific tag of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteTag.json
 */
async function apiManagementDeleteTag() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const tagId = "tagId1";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.delete(resourceGroupName, serviceName, tagId, ifMatch);
  console.log(result);
}

apiManagementDeleteTag().catch(console.error);
