const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates the IdentityProvider configuration.
 *
 * @summary Creates or Updates the IdentityProvider configuration.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateIdentityProvider.json
 */
async function apiManagementCreateIdentityProvider() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const identityProviderName = "facebook";
  const parameters = {
    clientId: "facebookid",
    clientSecret: "facebookapplicationsecret",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.identityProvider.createOrUpdate(
    resourceGroupName,
    serviceName,
    identityProviderName,
    parameters
  );
  console.log(result);
}
