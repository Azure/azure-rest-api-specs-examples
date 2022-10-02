const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing API Management service.
 *
 * @summary Updates an existing API Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateServiceToNewVnetAndAZs.json
 */
async function apiManagementUpdateServiceToNewVnetAndAvailabilityZones() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const parameters = {
    publicIpAddressId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/apimService1ip",
    sku: { name: "Premium", capacity: 2 },
    virtualNetworkConfiguration: {
      subnetResourceId:
        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimService1v2/subnets/default",
    },
    virtualNetworkType: "External",
    zones: ["1", "2"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}

apiManagementUpdateServiceToNewVnetAndAvailabilityZones().catch(console.error);
