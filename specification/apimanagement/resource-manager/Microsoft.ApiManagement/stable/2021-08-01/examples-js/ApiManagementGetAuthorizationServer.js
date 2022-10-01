const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the authorization server specified by its identifier.
 *
 * @summary Gets the details of the authorization server specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetAuthorizationServer.json
 */
async function apiManagementGetAuthorizationServer() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const authsid = "newauthServer2";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationServer.get(resourceGroupName, serviceName, authsid);
  console.log(result);
}

apiManagementGetAuthorizationServer().catch(console.error);
