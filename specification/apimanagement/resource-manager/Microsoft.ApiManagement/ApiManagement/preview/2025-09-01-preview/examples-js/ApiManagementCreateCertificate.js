const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the certificate being used for authentication with the backend.
 *
 * @summary creates or updates the certificate being used for authentication with the backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateCertificate.json
 */
async function apiManagementCreateCertificate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.certificate.createOrUpdate("rg1", "apimService1", "tempcert", {
    data: "****************Base 64 Encoded Certificate *******************************",
    password: "****Certificate Password******",
  });
  console.log(result);
}
