const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing IdentityProvider configuration.
 *
 * @summary Updates an existing IdentityProvider configuration.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateIdentityProvider.json
 */
async function apiManagementUpdateIdentityProvider() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const identityProviderName = "facebook";
  const ifMatch = "*";
  const parameters = {
    clientId: "updatedfacebookid",
    clientSecret: "updatedfacebooksecret",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.identityProvider.update(
    resourceGroupName,
    serviceName,
    identityProviderName,
    ifMatch,
    parameters
  );
  console.log(result);
}
