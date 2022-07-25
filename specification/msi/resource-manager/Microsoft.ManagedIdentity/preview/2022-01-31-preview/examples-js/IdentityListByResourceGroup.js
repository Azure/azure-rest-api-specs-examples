const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the userAssignedIdentities available under the specified ResourceGroup.
 *
 * @summary Lists all the userAssignedIdentities available under the specified ResourceGroup.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/IdentityListByResourceGroup.json
 */
async function identityListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgName";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.userAssignedIdentities.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

identityListByResourceGroup().catch(console.error);
