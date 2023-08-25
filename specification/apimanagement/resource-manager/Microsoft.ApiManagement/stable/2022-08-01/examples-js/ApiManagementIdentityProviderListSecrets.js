const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the client secret details of the Identity Provider.
 *
 * @summary Gets the client secret details of the Identity Provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementIdentityProviderListSecrets.json
 */
async function apiManagementIdentityProviderListSecrets() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const identityProviderName = "aadB2C";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.identityProvider.listSecrets(
    resourceGroupName,
    serviceName,
    identityProviderName
  );
  console.log(result);
}
