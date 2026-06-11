const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateBackendWithCustomAssignedCertificate_X509Names.json
 */
async function apiManagementCreateBackendCustomAssignedCertificateX509Names() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.createOrUpdate("rg1", "apimService1", "cabackend", {
    description: "Custom Assigned certificate Example",
    tls: {
      serverX509Names: [{ name: "*.badssl.com" }],
      validateCertificateChain: true,
      validateCertificateName: true,
    },
    url: "https://self-signed.badssl.com/",
    protocol: "http",
  });
  console.log(result);
}
