const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an identity in the specified subscription and resource group.
 *
 * @summary Update an identity in the specified subscription and resource group.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/IdentityUpdate.json
 */
async function identityUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const resourceName = "resourceName";
  const parameters = {
    location: "eastus",
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.userAssignedIdentities.update(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

identityUpdate().catch(console.error);
