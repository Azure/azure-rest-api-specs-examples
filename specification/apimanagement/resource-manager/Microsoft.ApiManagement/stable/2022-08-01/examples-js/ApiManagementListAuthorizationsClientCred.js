const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of authorization providers defined within a authorization provider.
 *
 * @summary Lists a collection of authorization providers defined within a authorization provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationsClientCred.json
 */
async function apiManagementListAuthorizationsClientCred() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithclientcred";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.authorization.listByAuthorizationProvider(
    resourceGroupName,
    serviceName,
    authorizationProviderId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
