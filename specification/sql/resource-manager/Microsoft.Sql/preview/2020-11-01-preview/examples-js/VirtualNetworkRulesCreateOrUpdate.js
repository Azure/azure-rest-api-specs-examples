const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an existing virtual network rule.
 *
 * @summary Creates or updates an existing virtual network rule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualNetworkRulesCreateOrUpdate.json
 */
async function createOrUpdateAVirtualNetworkRule() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default";
  const serverName = "vnet-test-svr";
  const virtualNetworkRuleName = "vnet-firewall-rule";
  const parameters = {
    ignoreMissingVnetServiceEndpoint: false,
    virtualNetworkSubnetId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    virtualNetworkRuleName,
    parameters,
  );
  console.log(result);
}
