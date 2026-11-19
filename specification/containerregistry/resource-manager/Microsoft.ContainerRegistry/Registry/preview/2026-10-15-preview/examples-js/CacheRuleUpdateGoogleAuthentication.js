const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a cache rule for a container registry with the specified parameters.
 *
 * @summary updates a cache rule for a container registry with the specified parameters.
 * x-ms-original-file: 2026-10-15-preview/CacheRuleUpdateGoogleAuthentication.json
 */
async function updateGoogleAuthenticationWhilePreservingOmittedFields() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.cacheRules.update("myResourceGroup", "myRegistry", "myCacheRule", {
    additionalAuthenticationProperties: {
      authenticationType: "GoogleArtifactRegistry",
      workloadIdentityProvider: "updated-identity-provider",
    },
  });
  console.log(result);
}
