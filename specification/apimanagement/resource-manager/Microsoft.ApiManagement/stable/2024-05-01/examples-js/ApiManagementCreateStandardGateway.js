const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an API Management gateway. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates or updates an API Management gateway. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateStandardGateway.json
 */
async function apiManagementCreateStandardGateway() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "apimGateway1";
  const parameters = {
    backend: {
      subnet: {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vn1/subnets/sn1",
      },
    },
    location: "South Central US",
    sku: { name: "Standard", capacity: 1 },
    tags: { name: "Contoso", test: "User" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiGateway.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gatewayName,
    parameters,
  );
  console.log(result);
}
