const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to From KeyVault, Refresh the certificate being used for authentication with the backend.
 *
 * @summary From KeyVault, Refresh the certificate being used for authentication with the backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementRefreshCertificate.json
 */
async function apiManagementRefreshCertificate() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const certificateId = "templateCertkv";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.certificate.refreshSecret(
    resourceGroupName,
    serviceName,
    certificateId
  );
  console.log(result);
}
