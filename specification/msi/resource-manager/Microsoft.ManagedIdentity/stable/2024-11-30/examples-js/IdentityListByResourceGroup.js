const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all the userAssignedIdentities available under the specified ResourceGroup.
 *
 * @summary Lists all the userAssignedIdentities available under the specified ResourceGroup.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/IdentityListByResourceGroup.json
 */
async function identityListByResourceGroup() {
  const subscriptionId = process.env["MSI_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MSI_RESOURCE_GROUP"] || "rgName";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.userAssignedIdentities.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
