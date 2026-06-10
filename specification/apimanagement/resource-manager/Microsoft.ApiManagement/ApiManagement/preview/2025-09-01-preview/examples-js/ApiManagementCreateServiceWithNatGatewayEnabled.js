const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateServiceWithNatGatewayEnabled.json
 */
async function apiManagementCreateServiceWithNatGatewayEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.createOrUpdate("rg1", "apimService1", {
    location: "East US",
    natGatewayState: "Enabled",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    sku: { name: "Premium", capacity: 1 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
  });
  console.log(result);
}
