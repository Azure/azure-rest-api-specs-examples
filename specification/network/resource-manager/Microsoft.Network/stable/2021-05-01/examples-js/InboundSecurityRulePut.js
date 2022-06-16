const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
 *
 * @summary Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/InboundSecurityRulePut.json
 */
async function createNetworkVirtualApplianceInboundSecurityRules() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const ruleCollectionName = "rule1";
  const parameters = {
    rules: [
      {
        destinationPortRange: 22,
        sourceAddressPrefix: "50.20.121.5/32",
        protocol: "TCP",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.inboundSecurityRuleOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkVirtualApplianceName,
    ruleCollectionName,
    parameters
  );
  console.log(result);
}

createNetworkVirtualApplianceInboundSecurityRules().catch(console.error);
