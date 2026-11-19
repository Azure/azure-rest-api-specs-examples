const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a cache rule for a container registry with the specified parameters.
 *
 * @summary creates a cache rule for a container registry with the specified parameters.
 * x-ms-original-file: 2026-10-15-preview/CacheRuleCreateUserAssignedMIEntraFederatedIdentityCredential.json
 */
async function cacheRuleCreateUserAssignedMIAuthEntraFederatedIdentityCredential() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.cacheRules.create("myResourceGroup", "myRegistry", "myCacheRule", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity":
          {},
      },
    },
    sourceRepository: "myupstream.azurecr.io/hello-world",
    targetRepository: "cached-upstream/hello-world",
    additionalAuthenticationProperties: {
      authenticationType: "EntraFederatedIdentityCredential",
      sourceTenantId: "72f988bf-86f1-41af-91ab-2d7cd011db47",
      federatedApplicationClientId: "abcdef12-3456-7890-abcd-ef1234567890",
    },
  });
  console.log(result);
}
