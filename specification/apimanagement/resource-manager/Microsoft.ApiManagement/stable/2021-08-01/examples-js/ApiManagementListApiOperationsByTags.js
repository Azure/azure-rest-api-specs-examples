const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of operations associated with tags.
 *
 * @summary Lists a collection of operations associated with tags.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiOperationsByTags.json
 */
async function apiManagementListApiOperationsByTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "a1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operationOperations.listByTags(
    resourceGroupName,
    serviceName,
    apiId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListApiOperationsByTags().catch(console.error);
