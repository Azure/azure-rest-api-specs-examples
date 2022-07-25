const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the identity.
 *
 * @summary Deletes the identity.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/IdentityDelete.json
 */
async function identityDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const resourceName = "resourceName";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.userAssignedIdentities.delete(resourceGroupName, resourceName);
  console.log(result);
}

identityDelete().catch(console.error);
