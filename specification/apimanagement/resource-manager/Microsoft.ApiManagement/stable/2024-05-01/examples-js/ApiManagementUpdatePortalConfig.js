const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update the developer portal configuration.
 *
 * @summary Update the developer portal configuration.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdatePortalConfig.json
 */
async function apiManagementUpdatePortalConfig() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const portalConfigId = "default";
  const ifMatch = "*";
  const parameters = {
    cors: { allowedOrigins: ["https://contoso.com"] },
    csp: {
      allowedSources: ["*.contoso.com"],
      mode: "reportOnly",
      reportUri: ["https://report.contoso.com"],
    },
    delegation: {
      delegateRegistration: false,
      delegateSubscription: false,
      delegationUrl: undefined,
      validationKey: undefined,
    },
    enableBasicAuth: true,
    signin: { require: false },
    signup: {
      termsOfService: {
        requireConsent: false,
        text: "I agree to the service terms and conditions.",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.portalConfig.update(
    resourceGroupName,
    serviceName,
    portalConfigId,
    ifMatch,
    parameters,
  );
  console.log(result);
}
