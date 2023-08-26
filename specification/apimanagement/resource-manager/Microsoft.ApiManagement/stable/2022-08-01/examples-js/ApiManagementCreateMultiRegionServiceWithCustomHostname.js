const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
 */
async function apiManagementCreateMultiRegionServiceWithCustomHostname() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = {
    additionalLocations: [
      {
        disableGateway: true,
        location: "East US",
        sku: { name: "Premium", capacity: 1 },
      },
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
    ],
    location: "West US",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    sku: { name: "Premium", capacity: 1 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
    virtualNetworkType: "None",
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
