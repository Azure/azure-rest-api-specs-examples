Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of inbound NAT rule port mappings.
 *
 * @summary List of inbound NAT rule port mappings.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/QueryInboundNatRulePortMapping.json
 */
async function queryInboundNatRulePortMapping() {
  const subscriptionId = "subid";
  const groupName = "rg1";
  const loadBalancerName = "lb1";
  const backendPoolName = "bp1";
  const parameters = {
    ipAddress: "10.0.0.4",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.beginListInboundNatRulePortMappingsAndWait(
    groupName,
    loadBalancerName,
    backendPoolName,
    parameters
  );
  console.log(result);
}

queryInboundNatRulePortMapping().catch(console.error);
```
