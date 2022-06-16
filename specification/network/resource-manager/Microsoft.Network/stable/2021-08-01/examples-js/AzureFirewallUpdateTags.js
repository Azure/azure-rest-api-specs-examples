const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags of an Azure Firewall resource.
 *
 * @summary Updates tags of an Azure Firewall resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/AzureFirewallUpdateTags.json
 */
async function updateAzureFirewallTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "azfwtest";
  const azureFirewallName = "fw1";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.beginUpdateTagsAndWait(
    resourceGroupName,
    azureFirewallName,
    parameters
  );
  console.log(result);
}

updateAzureFirewallTags().catch(console.error);
