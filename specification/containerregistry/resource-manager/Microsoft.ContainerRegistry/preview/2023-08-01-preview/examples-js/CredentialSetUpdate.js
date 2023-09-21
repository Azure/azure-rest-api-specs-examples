const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a credential set for a container registry with the specified parameters.
 *
 * @summary Updates a credential set for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/CredentialSetUpdate.json
 */
async function credentialSetUpdate() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const credentialSetName = "myCredentialSet";
  const credentialSetUpdateParameters = {
    authCredentials: [
      {
        name: "Credential1",
        passwordSecretIdentifier: "https://myvault.vault.azure.net/secrets/password2",
        usernameSecretIdentifier: "https://myvault.vault.azure.net/secrets/username2",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.credentialSets.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    credentialSetName,
    credentialSetUpdateParameters
  );
  console.log(result);
}
