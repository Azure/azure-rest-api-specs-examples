const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateServiceInVnetWithPublicIP.json
 */
async function apiManagementCreateServiceInVnetWithPublicIP() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.createOrUpdate("rg1", "apimService1", {
    location: "East US 2 EUAP",
    publicIpAddressId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgName/providers/Microsoft.Network/publicIPAddresses/apimazvnet",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    virtualNetworkConfiguration: {
      subnetResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgName/providers/Microsoft.Network/virtualNetworks/apimcus/subnets/tenant",
    },
    virtualNetworkType: "External",
    sku: { name: "Premium", capacity: 2 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
    zones: ["1", "2"],
  });
  console.log(result);
}
