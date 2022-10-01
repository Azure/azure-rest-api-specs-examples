const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the certificate specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the certificate specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadCertificate.json
 */
async function apiManagementHeadCertificate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const certificateId = "templateCert1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.certificate.getEntityTag(
    resourceGroupName,
    serviceName,
    certificateId
  );
  console.log(result);
}

apiManagementHeadCertificate().catch(console.error);
