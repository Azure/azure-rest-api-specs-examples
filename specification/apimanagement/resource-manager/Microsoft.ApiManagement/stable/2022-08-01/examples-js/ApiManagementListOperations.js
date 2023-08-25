const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations of the Microsoft.ApiManagement provider.
 *
 * @summary Lists all of the available REST API operations of the Microsoft.ApiManagement provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListOperations.json
 */
async function apiManagementListOperations() {
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.apiManagementOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
