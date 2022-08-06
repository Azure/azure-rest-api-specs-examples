const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists active security admin rules in a network manager.
 *
 * @summary Lists active security admin rules in a network manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerActiveSecurityAdminRulesList.json
 */
async function listActiveSecurityAdminRules() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const networkManagerName = "testNetworkManager";
  const parameters = {
    regions: ["westus"],
    skipToken: "fakeSkipTokenCode",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.listActiveSecurityAdminRules(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

listActiveSecurityAdminRules().catch(console.error);
