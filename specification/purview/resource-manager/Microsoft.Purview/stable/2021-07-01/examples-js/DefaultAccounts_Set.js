const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the default account for the scope.
 *
 * @summary Sets the default account for the scope.
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Set.json
 */
async function defaultAccountsSet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const defaultAccountPayload = {
    accountName: "myDefaultAccount",
    resourceGroupName: "rg-1",
    scope: "12345678-1234-1234-12345678abc",
    scopeTenantId: "12345678-1234-1234-12345678abc",
    scopeType: "Tenant",
    subscriptionId: "12345678-1234-1234-12345678aaa",
  };
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.defaultAccounts.set(defaultAccountPayload);
  console.log(result);
}

defaultAccountsSet().catch(console.error);
