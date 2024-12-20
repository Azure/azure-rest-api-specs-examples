const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the available specified Network Virtual Appliance Inbound Security Rules Collection.
 *
 * @summary Retrieves the available specified Network Virtual Appliance Inbound Security Rules Collection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/InboundSecurityRuleGet.json
 */
async function createNetworkVirtualApplianceInboundSecurityRules() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkVirtualApplianceName = "nva";
  const ruleCollectionName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.inboundSecurityRuleOperations.get(
    resourceGroupName,
    networkVirtualApplianceName,
    ruleCollectionName,
  );
  console.log(result);
}
