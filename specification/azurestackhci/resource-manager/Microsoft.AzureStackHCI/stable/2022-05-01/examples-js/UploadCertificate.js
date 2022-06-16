const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upload certificate.
 *
 * @summary Upload certificate.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/UploadCertificate.json
 */
async function uploadCertificate() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const uploadCertificateRequest = {
    properties: { certificates: ["base64cert", "base64cert"] },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusters.beginUploadCertificateAndWait(
    resourceGroupName,
    clusterName,
    uploadCertificateRequest
  );
  console.log(result);
}

uploadCertificate().catch(console.error);
