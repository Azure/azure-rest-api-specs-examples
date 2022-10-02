const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update Sign-In settings.
 *
 * @summary Update Sign-In settings.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsUpdateSignIn.json
 */
async function apiManagementPortalSettingsUpdateSignIn() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const ifMatch = "*";
  const parameters = { enabled: true };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.signInSettings.update(
    resourceGroupName,
    serviceName,
    ifMatch,
    parameters
  );
  console.log(result);
}

apiManagementPortalSettingsUpdateSignIn().catch(console.error);
