const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a cache rule for a container registry with the specified parameters.
 *
 * @summary creates a cache rule for a container registry with the specified parameters.
 * x-ms-original-file: 2026-03-01-preview/CacheRuleCreateUserAssignedMIGoogle.json
 */
async function cacheRuleCreateUserAssignedMIAuthGoogle() {
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
    sourceRepository: "us-west1-docker.pkg.dev/repository/hello-world",
    targetRepository: "cached-acr/hello-world",
    additionalAuthenticationProperties: {
      authenticationType: "GoogleArtifactRegistry",
      projectNumber: "123456789012",
      workloadIdentityPool: "my-workload-identity-pool",
      workloadIdentityProvider: "my-workload-identity-provider",
    },
  });
  console.log(result);
}
