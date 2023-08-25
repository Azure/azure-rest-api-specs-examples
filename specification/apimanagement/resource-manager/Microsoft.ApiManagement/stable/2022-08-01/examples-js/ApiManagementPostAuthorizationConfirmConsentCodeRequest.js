const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Confirm valid consent code to suppress Authorizations anti-phishing page.
 *
 * @summary Confirm valid consent code to suppress Authorizations anti-phishing page.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPostAuthorizationConfirmConsentCodeRequest.json
 */
async function apiManagementPostAuthorizationConfirmConsentCodeRequest() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithauthcode";
  const authorizationId = "authz1";
  const parameters = {
    consentCode: "theconsentcode",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorization.confirmConsentCode(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    authorizationId,
    parameters
  );
  console.log(result);
}
