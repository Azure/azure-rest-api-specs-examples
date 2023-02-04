const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the userAssignedIdentities available under the specified subscription.
 *
 * @summary Lists all the userAssignedIdentities available under the specified subscription.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityListBySubscription.json
 */
async function identityListBySubscription() {
  const subscriptionId = process.env["MSI_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.userAssignedIdentities.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
