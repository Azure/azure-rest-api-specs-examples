const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Tags associated with the API.
 *
 * @summary Lists all Tags associated with the API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiTags.json
 */
async function apiManagementListApiTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "57d2ef278aa04f0888cba3f3";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tag.listByApi(resourceGroupName, serviceName, apiId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListApiTags().catch(console.error);
