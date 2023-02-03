const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the federated identity credential.
 *
 * @summary Deletes the federated identity credential.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialDelete.json
 */
async function federatedIdentityCredentialDelete() {
  const subscriptionId =
    process.env["MSI_SUBSCRIPTION_ID"] || "c267c0e7-0a73-4789-9e17-d26aeb0904e5";
  const resourceGroupName = process.env["MSI_RESOURCE_GROUP"] || "rgName";
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
