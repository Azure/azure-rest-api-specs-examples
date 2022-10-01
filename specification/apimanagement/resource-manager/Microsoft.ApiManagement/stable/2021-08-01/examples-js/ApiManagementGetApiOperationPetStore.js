const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the API Operation specified by its identifier.
 *
 * @summary Gets the details of the API Operation specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationPetStore.json
 */
async function apiManagementGetApiOperationPetStore() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "swagger-petstore";
  const operationId = "loginUser";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiOperation.get(resourceGroupName, serviceName, apiId, operationId);
  console.log(result);
}

apiManagementGetApiOperationPetStore().catch(console.error);
