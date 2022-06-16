const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update NetworkRuleSet for a Namespace.
 *
 * @summary Create or update NetworkRuleSet for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/EHNetworkRuleSetCreate.json
 */
async function nameSpaceNetworkRuleSetCreate() {
  const subscriptionId = "Subscription";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-6019";
  const parameters = {
    defaultAction: "Deny",
    ipRules: [
      { action: "Allow", ipMask: "1.1.1.1" },
      { action: "Allow", ipMask: "1.1.1.2" },
      { action: "Allow", ipMask: "1.1.1.3" },
      { action: "Allow", ipMask: "1.1.1.4" },
      { action: "Allow", ipMask: "1.1.1.5" },
    ],
    virtualNetworkRules: [
      {
        ignoreMissingVnetServiceEndpoint: true,
        subnet: {
          id: "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet2",
        },
      },
      {
        ignoreMissingVnetServiceEndpoint: false,
        subnet: {
          id: "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet3",
        },
      },
      {
        ignoreMissingVnetServiceEndpoint: false,
        subnet: {
          id: "/subscriptions/subscriptionid/resourcegroups/resourcegroupid/providers/Microsoft.Network/virtualNetworks/myvn/subnets/subnet6",
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.createOrUpdateNetworkRuleSet(
    resourceGroupName,
    namespaceName,
    parameters
  );
  console.log(result);
}

nameSpaceNetworkRuleSetCreate().catch(console.error);
