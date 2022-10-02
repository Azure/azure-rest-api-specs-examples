const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the configuration details of the identity Provider configured in specified service instance.
 *
 * @summary Gets the configuration details of the identity Provider configured in specified service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetIdentityProvider.json
 */
async function apiManagementGetIdentityProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const identityProviderName = "aadB2C";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.identityProvider.get(
    resourceGroupName,
    serviceName,
    identityProviderName
  );
  console.log(result);
}

apiManagementGetIdentityProvider().catch(console.error);
