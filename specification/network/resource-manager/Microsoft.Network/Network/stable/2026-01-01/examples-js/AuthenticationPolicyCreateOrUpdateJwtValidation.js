const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an authentication policy with the specified name within a resource group.
 *
 * @summary creates or updates an authentication policy with the specified name within a resource group.
 * x-ms-original-file: 2026-01-01/AuthenticationPolicyCreateOrUpdateJwtValidation.json
 */
async function createsOrUpdatesAJWTValidationAuthenticationPolicyWithinAResourceGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.authenticationPolicies.createOrUpdate("rg1", "jwtValidationPolicy", {
    location: "eastus",
    properties: {
      userTrustProviderType: "entra",
      onUnauthenticatedRequest: "deny",
      authenticationProperties: {
        issuer: "https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/",
        clientId: "00000000-0000-0000-0000-000000000001",
        jwksUri:
          "https://login.microsoftonline.com/11111111-1111-1111-1111-111111111111/discovery/v2.0/keys",
        audience: "api://myapp",
      },
    },
  });
  console.log(result);
}
