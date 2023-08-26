const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceWithSystemCertificates.json
 */
async function apiManagementCreateServiceWithSystemCertificates() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = {
    certificates: [
      {
        certificatePassword: "Password",
        encodedCertificate: "*******Base64 encoded Certificate******************",
        storeName: "CertificateAuthority",
      },
    ],
    location: "Central US",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    sku: { name: "Basic", capacity: 1 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}
