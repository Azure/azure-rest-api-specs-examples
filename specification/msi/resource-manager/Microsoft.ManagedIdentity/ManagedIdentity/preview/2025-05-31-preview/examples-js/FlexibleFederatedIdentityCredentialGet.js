const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the federated identity credential.
 *
 * @summary gets the federated identity credential.
 * x-ms-original-file: 2025-05-31-preview/FlexibleFederatedIdentityCredentialGet.json
 */
async function flexibleFederatedIdentityCredentialGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c267c0e7-0a73-4789-9e17-d26aeb0904e5";
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.federatedIdentityCredentials.get(
    "rgName",
    "resourceName",
    "ficResourceName",
  );
  console.log(result);
}
