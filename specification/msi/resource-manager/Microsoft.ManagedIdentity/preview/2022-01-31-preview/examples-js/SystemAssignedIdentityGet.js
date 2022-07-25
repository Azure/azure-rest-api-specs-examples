const { ManagedServiceIdentityClient } = require("@azure/arm-msi");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the systemAssignedIdentity available under the specified RP scope.
 *
 * @summary Gets the systemAssignedIdentity available under the specified RP scope.
 * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/SystemAssignedIdentityGet.json
 */
async function msiOperationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "scope";
  const credential = new DefaultAzureCredential();
  const client = new ManagedServiceIdentityClient(credential, subscriptionId);
  const result = await client.systemAssignedIdentities.getByScope(scope);
  console.log(result);
}

msiOperationsList().catch(console.error);
