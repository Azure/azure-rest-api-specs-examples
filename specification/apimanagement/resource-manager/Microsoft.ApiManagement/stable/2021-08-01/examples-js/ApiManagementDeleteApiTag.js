const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detach the tag from the Api.
 *
 * @summary Detach the tag from the Api.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiTag.json
 */
async function apiManagementDeleteApiTag() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "59d5b28d1f7fab116c282650";
  const tagId = "59d5b28e1f7fab116402044e";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.detachFromApi(resourceGroupName, serviceName, apiId, tagId);
  console.log(result);
}

apiManagementDeleteApiTag().catch(console.error);
