const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the federated identity credential.
 *
 * @summary Gets the federated identity credential.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/FederatedIdentityCredentialGet.json
 */
async function federatedIdentityCredentialGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const resourceName = "resourceName";
  const federatedIdentityCredentialResourceName = "ficResourceName";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.federatedIdentityCredentials.get(
    resourceGroupName,
    resourceName,
    federatedIdentityCredentialResourceName
  );
  console.log(result);
}

federatedIdentityCredentialGet().catch(console.error);
