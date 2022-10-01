const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Sign Up Settings for the Portal
 *
 * @summary Get Sign Up Settings for the Portal
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsGetSignUp.json
 */
async function apiManagementPortalSettingsGetSignUp() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.signUpSettings.get(resourceGroupName, serviceName);
  console.log(result);
}

apiManagementPortalSettingsGetSignUp().catch(console.error);
