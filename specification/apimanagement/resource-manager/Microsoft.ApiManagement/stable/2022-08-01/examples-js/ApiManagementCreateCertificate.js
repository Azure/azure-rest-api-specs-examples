const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the certificate being used for authentication with the backend.
 *
 * @summary Creates or updates the certificate being used for authentication with the backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateCertificate.json
 */
async function apiManagementCreateCertificate() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const certificateId = "tempcert";
  const parameters = {
    data: "****************Base 64 Encoded Certificate *******************************",
    password: "****Certificate Password******",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.certificate.createOrUpdate(
    resourceGroupName,
    serviceName,
    certificateId,
    parameters
  );
  console.log(result);
}
