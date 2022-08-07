const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a load balancer inbound NAT rule.
 *
 * @summary Creates or updates a load balancer inbound NAT rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/InboundNatRuleCreate.json
 */
async function inboundNatRuleCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb1";
  const inboundNatRuleName = "natRule1.1";
  const inboundNatRuleParameters = {
    backendPort: 3389,
    enableFloatingIP: false,
    enableTcpReset: false,
    frontendIPConfiguration: {
      id: "/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1",
    },
    frontendPort: 3390,
    idleTimeoutInMinutes: 4,
    protocol: "Tcp",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.inboundNatRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    loadBalancerName,
    inboundNatRuleName,
    inboundNatRuleParameters
  );
  console.log(result);
}

inboundNatRuleCreate().catch(console.error);
