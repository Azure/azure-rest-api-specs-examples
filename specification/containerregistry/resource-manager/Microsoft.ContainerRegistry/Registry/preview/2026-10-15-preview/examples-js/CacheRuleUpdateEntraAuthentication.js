const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a cache rule for a container registry with the specified parameters.
 *
 * @summary updates a cache rule for a container registry with the specified parameters.
 * x-ms-original-file: 2026-10-15-preview/CacheRuleUpdateEntraAuthentication.json
 */
async function updateEntraAuthenticationWhilePreservingOmittedFields() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.cacheRules.update("myResourceGroup", "myRegistry", "myCacheRule", {
    additionalAuthenticationProperties: {
      authenticationType: "EntraFederatedIdentityCredential",
      federatedApplicationClientId: "fedcba98-7654-3210-abcd-ef1234567890",
    },
  });
  console.log(result);
}
