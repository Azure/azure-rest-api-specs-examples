const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the federated identity credential.
 *
 * @summary Deletes the federated identity credential.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/FederatedIdentityCredentialDelete.json
 */
async function federatedIdentityCredentialDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const resourceName = "resourceName";
  const federatedIdentityCredentialResourceName = "ficResourceName";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.federatedIdentityCredentials.delete(
    resourceGroupName,
    resourceName,
    federatedIdentityCredentialResourceName
  );
  console.log(result);
}

federatedIdentityCredentialDelete().catch(console.error);
