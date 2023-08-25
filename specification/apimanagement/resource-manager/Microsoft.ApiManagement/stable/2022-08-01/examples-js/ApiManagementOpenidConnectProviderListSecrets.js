const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the client secret details of the OpenID Connect Provider.
 *
 * @summary Gets the client secret details of the OpenID Connect Provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementOpenidConnectProviderListSecrets.json
 */
async function apiManagementOpenidConnectProviderListSecrets() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const opid = "templateOpenIdConnect2";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.openIdConnectProvider.listSecrets(
    resourceGroupName,
    serviceName,
    opid
  );
  console.log(result);
}
