const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing API Management service.
 *
 * @summary updates an existing API Management service.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementUpdateServiceDisableTls10.json
 */
async function apiManagementUpdateServiceDisableTls10() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.update("rg1", "apimService1", {
    customProperties: {
      "Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": "false",
    },
  });
  console.log(result);
}
