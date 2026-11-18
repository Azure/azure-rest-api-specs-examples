const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a cache rule for a container registry with the specified parameters.
 *
 * @summary creates a cache rule for a container registry with the specified parameters.
 * x-ms-original-file: 2026-10-15-preview/CacheRuleCreateUserAssignedMIAws.json
 */
async function cacheRuleCreateUserAssignedMIAuthAws() {
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
    sourceRepository: "123456789012.dkr.ecr.us-west-2.amazonaws.com/hello-world",
    targetRepository: "cached-ecr/hello-world",
    additionalAuthenticationProperties: {
      authenticationType: "AwsElasticContainerRegistry",
      accountNumber: "123456789012",
      region: "us-west-2",
      roleName: "my-ecr-pull-role",
      rolePath: "/acr/",
    },
  });
  console.log(result);
}
