const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
 */
async function apiManagementCreateMultiRegionServiceWithCustomHostname() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.createOrUpdate("rg1", "apimService1", {
    location: "West US",
    additionalLocations: [
      { disableGateway: true, location: "East US", sku: { name: "Premium", capacity: 1 } },
    ],
    apiVersionConstraint: { minApiVersion: "2019-01-01" },
    hostnameConfigurations: [
      {
        type: "Proxy",
        certificatePassword: "Password",
        defaultSslBinding: true,
        encodedCertificate: "****** Base 64 Encoded Certificate ************",
        hostName: "gateway1.msitesting.net",
      },
      {
        type: "Management",
        certificatePassword: "Password",
        encodedCertificate: "****** Base 64 Encoded Certificate ************",
        hostName: "mgmt.msitesting.net",
      },
      {
        type: "Portal",
        certificatePassword: "Password",
        encodedCertificate: "****** Base 64 Encoded Certificate ************",
        hostName: "portal1.msitesting.net",
      },
      {
        type: "ConfigurationApi",
        certificatePassword: "Password",
        encodedCertificate: "****** Base 64 Encoded Certificate ************",
        hostName: "configuration-api.msitesting.net",
      },
    ],
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    virtualNetworkType: "None",
    sku: { name: "Premium", capacity: 1 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
  });
  console.log(result);
}
