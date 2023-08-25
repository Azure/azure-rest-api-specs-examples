const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update Sign-Up settings.
 *
 * @summary Update Sign-Up settings.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPortalSettingsUpdateSignUp.json
 */
async function apiManagementPortalSettingsUpdateSignUp() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const ifMatch = "*";
  const parameters = {
    enabled: true,
    termsOfService: {
      consentRequired: true,
      enabled: true,
      text: "Terms of service text.",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.signUpSettings.update(
    resourceGroupName,
    serviceName,
    ifMatch,
    parameters
  );
  console.log(result);
}
