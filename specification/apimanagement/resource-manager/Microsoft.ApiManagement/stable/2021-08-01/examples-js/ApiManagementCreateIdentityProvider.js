const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates the IdentityProvider configuration.
 *
 * @summary Creates or Updates the IdentityProvider configuration.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateIdentityProvider.json
 */
async function apiManagementCreateIdentityProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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

apiManagementCreateIdentityProvider().catch(console.error);
